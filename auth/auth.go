package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt"
	"github.com/husio/plop/secret"
	"golang.org/x/net/context"
)

func Authenticated(ctx context.Context, r *http.Request) (*Token, bool) {
	auth := r.Header.Get("Authorization")
	chunks := strings.SplitN(auth, " ", 2)
	if len(chunks) != 2 {
		return nil, false
	}
	token, err := jwt.Parse(chunks[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(secret.SigKey), nil
	})
	if err != nil || !token.Valid {
		log.Printf("rejecting token: %+v: %s", token, err)
		return nil, false
	}
	tk := Token{
		Role:   token.Claims["role"].(string),
		UserID: token.Claims["uid"].(string),
	}
	return &tk, true
}

type Token struct {
	Role   string
	UserID string
}
