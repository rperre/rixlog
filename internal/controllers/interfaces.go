package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strings"
	"text/template"
)

func HttpResponse(w http.ResponseWriter, r *http.Request, response render.Renderer) {
	accpet_header := strings.Split(r.Header["Accept"][0], ",")[0]
	switch accpet_header {
	case "text/html":
		render.Status(r, http.StatusOK)
		tmpl := template.Must(template.ParseFiles("internal/views/article.html"))
		err := tmpl.Execute(w, response)
		if err != nil {
			errorData := make(map[string]string)
			errorData["message"] = err.Error()
			HttpErrorJSONResponse(w, r, HttpError{Data: errorData, Code: http.StatusUnprocessableEntity})
		}
	case "application/json":
		render.Status(r, http.StatusOK)
		err := render.Render(w, r, response)

		if err != nil {
			errorData := make(map[string]string)
			errorData["message"] = err.Error()
			HttpErrorJSONResponse(w, r, HttpError{Data: errorData, Code: http.StatusUnprocessableEntity})
		}
	default:
		errorData := make(map[string]string)
		errorData["message"] = "Content-Type not supported"
		HttpErrorJSONResponse(w, r, HttpError{Data: errorData, Code: http.StatusBadRequest})
	}
}

func (a *ArticlesController) Void(w http.ResponseWriter, r *http.Request) {}

type HttpError struct {
	Data map[string]string
	Code int
}

func (rd *HttpError) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HttpErrorJSONResponse(w http.ResponseWriter, r *http.Request, err HttpError) {
	render.Status(r, err.Code)
	render.JSON(w, r, err)
}

type ControllerRoute interface {
	Routes() chi.Router
}

type RouteMap map[string]ControllerRoute

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
