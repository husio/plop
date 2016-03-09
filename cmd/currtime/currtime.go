package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/husio/plop/auth"
	"github.com/husio/plop/discovery"
	"github.com/husio/web"
	"golang.org/x/net/context"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ltime)

	ctx := context.Background()
	app := application{
		ctx: ctx,
		rt: web.NewRouter("", web.Routes{
			web.GET(`/`, "", handleCurrentTime),
			web.ANY(`.*`, "", handleNotFound),
		}),
	}

	httpAddr := discovery.Any("currtime")
	log.Printf("running HTTP: %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, &app); err != nil {
		log.Fatalf("HTTP server error: %s", err)
	}
}

func handleCurrentTime(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token, ok := auth.Authenticated(ctx, r)
	if !ok {
		web.StdJSONErr(w, http.StatusUnauthorized)
		return
	}
	if token.Role != "service" {
		// just to differ from not authenticated
		web.StdJSONErr(w, http.StatusForbidden)
		return
	}

	resp := struct {
		Now time.Time
	}{
		Now: time.Now(),
	}
	web.JSONResp(w, resp, http.StatusOK)
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
