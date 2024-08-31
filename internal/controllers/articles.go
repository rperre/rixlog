package controllers

import (
	"context"
	"net/http"
	"rixlog/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ArticlesController struct{}

func (a *ArticlesController) Routes() chi.Router {
	r := chi.NewRouter()
	r.
		With(ExistingArticle).
		Get("/{id}", a.GetArticle)
	r.
		With(Authenticated).
		Post("/", a.Void)
	r.
		With(Authenticated).
		With(ExistingArticle).
		Put("/{id}", a.Void)
	r.
		With(Authenticated).
		With(AdminOnly).
		With(ExistingArticle).
		Delete("/{id}", a.Void)
	return r
}

func ExistingArticle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		model := &models.Article{}
		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if article, err := model.GetByID(id); err != nil {
			errorData := make(map[string]string)
			errorData["message"] = err.Error()
			HttpErrorJSONResponse(w, r, HttpError{Data: errorData, Code: http.StatusUnprocessableEntity})
		} else {
			ctx := context.WithValue(r.Context(), "article", article)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func (a *ArticlesController) GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*models.Article)
	if !ok {
		errorData := make(map[string]string)
		errorData["message"] = "Something went wrong"
		HttpErrorJSONResponse(w, r, HttpError{Data: errorData, Code: http.StatusUnprocessableEntity})
	}
	HttpResponse(w, r, article)
}
