package api

import (
	"context"
	"net/http"
)

func isValidAPIKey(key string) bool {
	return key == "api123"
}

func WithAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if !isValidAPIKey(key) {
			respondErr(w, r, http.StatusUnauthorized, "invalid API key")
			return
		}
		ctx := context.WithValue(r.Context(), ContextKeyAPIKey, key)
		fn(w, r.WithContext(ctx))
	}
}

func WithCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers",
			"Location")
		fn(w, r)
	}
}
