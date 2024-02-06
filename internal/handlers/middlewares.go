package handlers

import (
	"golang.org/x/net/context"
	"net/http"
)

const (
	KeyUserID = "userID"
)

func (h *Handler) AdminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	code before Handler
		value := 0
		ctx := context.WithValue(r.Context(), KeyUserID, value)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		//	code after Handler
	})
}

func (h *Handler) ClientAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//code before Handler
		userID := 0
		ctx := context.WithValue(r.Context(), KeyUserID, userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		//code after Handler
	})
}
