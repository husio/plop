package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt"
	"github.com/husio/plop/discovery"
	"github.com/husio/plop/secret"
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
			web.POST(`/`, "", handleLogin),
			web.ANY(`.*`, "", handleNotFound),
		}),
	}

	httpAddr := discovery.Any("auth")
	log.Printf("running HTTP: %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, &app); err != nil {
		log.Fatalf("HTTP server error: %s", err)
	}
}

func handleLogin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var input struct {
		Login    string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		web.JSONErr(w, err.Error(), http.StatusBadRequest)
		return
	}
	role := secret.AuthRole(input.Login, input.Password)
	if role == "" {
		web.JSONErr(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	user := User{
		ID:   fmt.Sprint(time.Now().Unix()), // because we don't store users
		Role: role,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	token.Claims["role"] = user.Role
	token.Claims["uid"] = user.ID
	raw, err := token.SignedString([]byte(secret.SigKey))
	if err != nil {
		msg := fmt.Sprintf("Cannot create JWT token: %s", err)
		web.JSONErr(w, msg, http.StatusInternalServerError)
		return
	}

	resp := struct {
		User  User
		Token string
	}{
		User:  user,
		Token: raw,
	}
	web.JSONResp(w, resp, http.StatusOK)
}

type User struct {
	ID   string
	Role string
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
