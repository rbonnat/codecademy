package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"

	"github.com/rbonnat/codecademy/user"
)

// Authorize is a middleware function that handles authorization
func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, claims, errToken := jwtauth.FromContext(r.Context())
		if errToken != nil {
			log.Printf("Error while getting Token from context: '%v'", errToken)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if !isAuthorized(claims, r.Method) {
			log.Printf("User is not authorized for this method: %s \n", r.Method)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// User is authorized
		next.ServeHTTP(w, r)
	})
}

func isAuthorized(claims jwt.MapClaims, method string) bool {
	authorized := false

	authorizations := claims[user.Authorization]

	switch method {
	case http.MethodGet:
		if r, exist := authorizations.(map[string]interface{})[user.Read]; exist {
			authorized = r.(bool)
		}
	case http.MethodDelete:
		if r, exist := authorizations.(map[string]interface{})[user.Delete]; exist {
			authorized = r.(bool)
		}
	case http.MethodPost:
		if r, exist := authorizations.(map[string]interface{})[user.Insert]; exist {
			authorized = r.(bool)
		}
	case http.MethodPut:
		if r, exist := authorizations.(map[string]interface{})[user.Update]; exist {
			authorized = r.(bool)
		}
	}

	return authorized
}
