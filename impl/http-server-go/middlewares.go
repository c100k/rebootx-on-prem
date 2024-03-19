package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"openapi"
)

const AUTHORIZATION_HEADER = "Authorization"
const CONTENT_TYPE = "application/json"

func authMiddleware(config *Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get(AUTHORIZATION_HEADER)

			if authorization != config.apiKey {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(openapi.NewErrorRes(Err401))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func headerMiddleware(config *Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", CONTENT_TYPE)

			next.ServeHTTP(w, r)
		})
	}
}

func logMiddleware(config *Config, logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL))

			next.ServeHTTP(w, r)
		})
	}
}
