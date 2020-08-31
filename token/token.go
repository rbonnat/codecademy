package token

import (
	"log"
	"os"

	"github.com/go-chi/jwtauth"

	"github.com/rbonnat/codecademy/secretstore"
)

const TokenSecretKey = "authorization/secret-key"
const tokenAlg = "HS256"

type Token struct {
	tokenAuth *jwtauth.JWTAuth
}

func New(secretStore *secretstore.SecretStore) *Token {
	secretKey, err := secretStore.Get(TokenSecretKey)
	if err != nil {
		log.Printf("Error while loading token secret: %v", err)
		os.Exit(1)
	}

	t := Token{
		tokenAuth: jwtauth.New(tokenAlg, []byte(secretKey), nil),
	}

	return &t
}

func (ts *Token) TokenAuth() *jwtauth.JWTAuth {
	return ts.tokenAuth
}
