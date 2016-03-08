package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt"
	"github.com/husio/envconf"
	"github.com/husio/web"
	"golang.org/x/net/context"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ltime)

	conf := struct {
		HTTP         string
		JWTSignature string
	}{
		HTTP: "localhost:8005",
	}
	envconf.Must(envconf.LoadEnv(&conf))

	ctx := context.Background()
	ctx = context.WithValue(ctx, "jwt:signature", conf.JWTSignature)
	app := application{
		ctx: ctx,
		rt: web.NewRouter("", web.Routes{
			web.GET(`/login`, "", handleLogin),
		}),
	}

	if err := http.ListenAndServe(conf.HTTP, &app); err != nil {
		log.Fatalf("HTTP server error: %s", err)
	}
}

func handleLogin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// authenticate everyone!
	user := User{
		ID:    fmt.Sprint(time.Now().Unix()),
		Admin: true,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	token.Claims["admin"] = user.Admin
	token.Claims["uid"] = user.ID
	raw, err := token.SignedString(jwtSigKey(ctx))
	if err != nil {
		msg := fmt.Sprintf("cannot create JWT token: %s", err)
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
	ID    string
	Admin bool
}

type application struct {
	ctx context.Context
	rt  *web.Router
}

func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.rt.ServeCtxHTTP(app.ctx, w, r)
}

func jwtSigKey(ctx context.Context) string {
	sig := ctx.Value("jwt:signature")
	if sig == nil {
		panic("JWT signature not in context")
	}
	return sig.(string)
}
