package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ControllerRoute interface {
	Routes() chi.Router
}

type ControllerMap map[string]ControllerRoute

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

// TODO: Create users, login, and a way to set admin flag
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		// perm := ctx.Value("acl.permission")
		// if !ok || !perm.IsAdmin() {
		// 	http.Error(w, http.StatusText(403), 403)
		// 	return
		// }

		fmt.Println("This must be for admin only!")

		http.Error(w, http.StatusText(401), 401)

		// next.ServeHTTP(w, r)
	})
}

// TODO: Create users, login, and a way to set admin flag
func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		// perm := ctx.Value("acl.permission")
		// if !ok || !perm.IsAdmin() {
		// 	http.Error(w, http.StatusText(403), 403)
		// 	return
		// }

		fmt.Println("This must be for authenticated only!")

		http.Error(w, http.StatusText(401), 401)

		// next.ServeHTTP(w, r)
	})
}
