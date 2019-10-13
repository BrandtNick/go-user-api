package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/brandtnick/rest/models"
	"github.com/dgrijalva/jwt-go"
)

// Authorization - JWT Authorization middleware
var Authorization = func(next http.Handler) http.Handler {

	auth := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authorize")

		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			res := map[string]interface{}{"message": "missing token"}
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")

			json.NewEncoder(w).Encode(res)
			return
		}

		t := strings.Split(bearerToken, " ")[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(t, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_secret")), nil
		})

		if err != nil {
			res := map[string]interface{}{"message": "Bad token"}
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")

			json.NewEncoder(w).Encode(res)
			return
		}

		if !token.Valid {
			res := map[string]interface{}{"message": "Invalid token"}
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")

			json.NewEncoder(w).Encode(res)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(auth)
}
