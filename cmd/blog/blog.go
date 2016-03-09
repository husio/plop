package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/husio/plop/auth"
	"github.com/husio/plop/discovery"
	"github.com/husio/plop/secret"
	"github.com/husio/web"
	"golang.org/x/net/context"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx := context.Background()
	ctx = WithDB(ctx, NewDatabase())

	app := application{
		ctx: ctx,
		rt: web.NewRouter("", web.Routes{
			web.GET(`/`, "", handleListEntries),
			web.POST(`/`, "", handleCreateEntry),
			web.ANY(`.*`, "", handleNotFound),
		}),
	}

	httpAddr := discovery.Any("blog")
	log.Printf("running HTTP: %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, &app); err != nil {
		log.Fatalf("HTTP server error: %s", err)
	}
}

func handleListEntries(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	db := DB(ctx)
	entries := db.List()
	resp := struct {
		Entries []Entry
	}{
		Entries: entries,
	}
	web.JSONResp(w, resp, http.StatusOK)
}

func handleCreateEntry(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// authenticate first

	var input struct {
		Title   string
		Content string
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		web.JSONErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, ok := auth.Authenticated(ctx, r)
	if !ok {
		web.StdJSONErr(w, http.StatusUnauthorized)
		return
	}

	var currtime struct {
		Now time.Time
	}
	if resp, err := GETService("currtime", "/"); err != nil {
		web.JSONErr(w, err.Error(), http.StatusInternalServerError)
		return
	} else if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		web.JSONErr(w, string(b), http.StatusInternalServerError)
		return
	} else {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&currtime); err != nil {
			web.JSONErr(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	db := DB(ctx)
	entry := db.Create(Entry{
		Created: currtime.Now,
		UserID:  token.UserID,
		Title:   input.Title,
		Content: input.Content,
	})
	web.JSONResp(w, entry, http.StatusCreated)
}

type application struct {
	ctx context.Context
	rt  *web.Router
}

func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.rt.ServeCtxHTTP(app.ctx, w, r)
}

func handleNotFound(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	web.StdJSONErr(w, http.StatusNotFound)
}

// GETService makes GET request to given service.
func GETService(service, path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://"+discovery.Any(service)+path, nil)
	if err != nil {
		return nil, err
	}

	token.mu.Lock()
	if token.raw == "" {
		var body bytes.Buffer
		if err := json.NewEncoder(&body).Encode(secret.Blog); err != nil {
			log.Printf("cannot serialize auth credentials: %s", err)
			return nil, err
		}
		resp, err := http.Post("http://"+discovery.Any("auth")+"/", "application/json", &body)
		if err != nil {
			log.Printf("token request failed: %s", err)
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("cannot authorize internally: %s", err)
		}
		var result struct {
			Token string
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("cannot decode auth response: %s", err)
		}
		token.raw = result.Token
	}
	defer token.mu.Unlock()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.raw)
	return http.DefaultClient.Do(req)
}

var token struct {
	mu  sync.Mutex
	raw string
}
