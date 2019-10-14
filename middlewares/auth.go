package middlewares

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/brandtnick/rest/models"
	"github.com/dgrijalva/jwt-go"
)

// Authorization - JWT Authorization middleware
var Authorization = func(next http.Handler) http.Handler {

	checkToken := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			res := map[string]interface{}{"response": "missing token"}
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")

			json.NewEncoder(w).Encode(res)
			return
		}

		authToken := strings.Split(authHeader, " ")[1]
		token, err := jwt.ParseWithClaims(authToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})
		if err != nil {
			res := map[string]interface{}{"response": "Bad token"}
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")

			json.NewEncoder(w).Encode(res)
			return
		}

		if !token.Valid {
			res := map[string]interface{}{"response": "Invalid token"}
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")

			json.NewEncoder(w).Encode(res)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(checkToken)
}
