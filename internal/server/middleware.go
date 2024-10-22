package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			slog.Info(r.RemoteAddr, r.Method, r.URL.Path)

			next(w, r)
		}
	}
}

func Auth() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")
			var res Response
			var login LoginData

			// Split the "Bearer" prefix and the token
			tokenParts := strings.SplitN(auth, " ", 2)
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				res = Response{Message: "Invalid Authorization header format"}
				json.NewEncoder(w).Encode(res)
				return
			}

			login, err := parseToken(tokenParts[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				res = Response{Message: "Invalid token"}
				json.NewEncoder(w).Encode(res)
				return
			}

			res = Response{Message: "Authorized...", Data: login}

			json.NewEncoder(os.Stdout).Encode(res)

			next(w, r)

		}

	}
}
