package middlewares

import (
	"log"
	"net/http"

	"github.com/fatih/color"
)

// LogHandler - basic http log from docs
func LogHandler(next http.Handler) http.Handler {
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(blue(r.Method) + " :: " + cyan(r.RequestURI))
		next.ServeHTTP(w, r)
	})
}
