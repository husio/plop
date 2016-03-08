package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt"
	"github.com/husio/envconf"
	"github.com/husio/web"
	"golang.org/x/net/context"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ltime)

	conf := struct {
		HTTP            string
		JWTSignatureKey string
	}{
		HTTP: "localhost:8005",
	}
	envconf.Must(envconf.LoadEnv(&conf))

	ctx := context.Background()
	ctx = WithDB(ctx, NewDatabase())
	ctx = context.WithValue(ctx, "jwt:signature", conf.JWTSignatureKey)

	app := application{
		ctx: ctx,
		rt: web.NewRouter("", web.Routes{
			web.GET(`/`, "", handleListEntries),
			web.POST(`/`, "", handleCreateEntry),
		}),
	}

	if err := http.ListenAndServe(conf.HTTP, &app); err != nil {
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

	token, ok := authenticated(ctx, r)
	if !ok {
		web.StdJSONErr(w, http.StatusUnauthorized)
		return
	}

	db := DB(ctx)
	entry := db.Create(Entry{
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

func authenticated(ctx context.Context, r *http.Request) (*Token, bool) {
	auth := r.Header.Get("Authorization")
	raw := strings.SplitN(auth, " ", 2)[1]
	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return sigPubKey(ctx), nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}
	tk := Token{
		Admin:  token.Claims["admin"].(bool),
		UserID: token.Claims["uid"].(string),
	}
	return &tk, true
}

type Token struct {
	Admin  bool
	UserID string
}

func sigPubKey(ctx context.Context) string {
	sig := ctx.Value("jwt:signature")
	if sig == nil {
		panic("JWT signature not in context")
	}
	return sig.(string)
}
