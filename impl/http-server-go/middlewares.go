package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"

	"openapi"
)

const AUTHORIZATION_HEADER = "Authorization"
const CONTENT_TYPE = "application/json"

func authMiddleware(config *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get(AUTHORIZATION_HEADER)

			if authorization != config.ApiKey {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(openapi.NewErrorRes(utils.Err401))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func headerMiddleware(_ *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", CONTENT_TYPE)

			next.ServeHTTP(w, r)
		})
	}
}

func logMiddleware(_ *config.Config, logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL))

			next.ServeHTTP(w, r)
		})
	}
}
