package middleware

import (
	"log"
	"net/http"
)

func printURl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.URL)
		next.ServeHTTP(w, r)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
	})
}
