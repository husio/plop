package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	etcd "github.com/coreos/etcd/client"
	"github.com/husio/envconf"
	"golang.org/x/net/context"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[regrun] ")
	log.SetFlags(log.Ltime | log.Lshortfile)

	if len(os.Args) < 2 {
		fail(2, "Usage: %s <host app> [<args>...]\n", os.Args[0])
	}

	conf := struct {
		ServiceID     string
		ServiceName   string
		ServiceHost   string
		ServicePort   int
		EtcdEndpoints []string
	}{
		ServiceID:     genID(4),
		ServiceName:   os.Args[1],
		EtcdEndpoints: []string{"http://localhost:2379"},
		ServiceHost:   "localhost",
		ServicePort:   randomPort(),
	}
	envconf.Must(envconf.LoadEnv(&conf))

	if conf.ServiceName == "" {
		fail(2, "SERVICE_NAME is not set\n")
	}

	ctx := context.Background()
	ctx, stop := context.WithCancel(ctx)
	defer stop()

	service := ServiceInfo{
		ID:   conf.ServiceID,
		Name: conf.ServiceName,
		Host: conf.ServiceHost,
		Port: conf.ServicePort,
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...) // todo what if there are no args?
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	env = append(env,
		fmt.Sprintf("SERVICE_HOST=%s", service.Host),
		fmt.Sprintf("SERVICE_PORT=%d", service.Port),
		fmt.Sprintf("SERVICE_ID=%s", service.ID),
		fmt.Sprintf("SERVICE_NAME=%s", service.Name),
	)
	cmd.Env = env

	log.Printf("running: %v", os.Args[1:])
	if err := cmd.Start(); err != nil {
		fail(1, "cannot start suprocess: %s", err)
	}

	errc := make(chan error, 1)

	go func() {
		if err := putConf(ctx, service, conf.EtcdEndpoints); err != nil {
			errc <- err
		}
	}()

	go func() {
		if err := cmd.Wait(); err != nil {
			errc <- fmt.Errorf("subprocess failed: %s", err)
		} else {
			errc <- nil
		}
	}()

	if err := <-errc; err != nil {
		fail(1, err.Error())
	}
}

type ServiceInfo struct {
	ID   string
	Name string
	Host string
	Port int
}

// putConf constanty updated in etcd given service information with expiration
// 5 seconds.
func putConf(ctx context.Context, s ServiceInfo, endpoints []string) error {
	var body string
	if b, err := json.MarshalIndent(s, "", "\t"); err != nil {
		return err
	} else {
		body = string(b)
	}

	c, err := etcd.New(etcd.Config{
		Endpoints:               endpoints,
		Transport:               etcd.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	})
	if err != nil {
		return err
	}
	api := etcd.NewKeysAPI(c)

	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()
	path := "/services/" + s.ID

	defer api.Delete(context.Background(), path, nil)

	opts := &etcd.SetOptions{TTL: 10 * time.Second}
	if _, err := api.Set(context.Background(), path, body, opts); err != nil {
		return fmt.Errorf("cannot set etcd value: %s", err)
	}
	// Refresh set to true means a TTL value can be updated
	// without firing a watch or changing the node value. A
	// value must not be provided when refreshing a key.
	opts.Refresh = true

	for {
		select {
		case <-tick.C:
		case <-ctx.Done():
			return ctx.Err()
		}

		// refresh TTL only
		if _, err := api.Set(context.Background(), path, "", opts); err != nil {
			return fmt.Errorf("cannot set etcd value: %s", err)
		}
	}
}

func genID(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := base32.StdEncoding.EncodeToString(b)
	return s[:length]
}

// fail print error message and exit with given code
func fail(exitCode int, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(exitCode)
}

// randomPort return random port number
func randomPort() int {
	var b [2]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	return 8000 + int(b[0]) + int(b[1])
}
