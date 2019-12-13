package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"src/models"
	u "src/utils"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//prefix
		api := "/api/v1/"
		auth := api + "auth/"
		courses := api + "courses/"
		notAuth := []string{auth + "login", auth + "register", courses}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // Grab the token from header
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //`Bearer {token-body}`
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) //Proceed in the middleware chain!
	})
}
