package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Token(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		token := strings.Split(tokenStr, " ")[1]

		fmt.Println(token)

		next.ServeHTTP(w, r)
	})
}
