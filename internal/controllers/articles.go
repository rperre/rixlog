package controllers

import (
	"encoding/json"
	"net/http"
	"rixlog/internal/databases"
	"rixlog/internal/models"

	"github.com/go-chi/chi/v5"
)

func Articles() *ArticlesController {
	if _Articles != nil {
		return _Articles
	}

	_Articles = &ArticlesController{}
	return _Articles
}

type ArticlesController struct{}

var _Articles *ArticlesController

func (a *ArticlesController) Routes() chi.Router {
	r := chi.NewRouter()
	r.With(paginate).Get("/", a.Sample)
	r.Get("/{slug}", a.Sample)
	r.With(Authenticated).Post("/", a.Sample)
	r.With(Authenticated).Put("/{slug}", a.Sample)
	r.With(Authenticated).With(AdminOnly).Delete("/{slug}", a.Sample)
	return r
}

func (a *ArticlesController) HandleDBHealth(w http.ResponseWriter, r *http.Request) {
	Sqlite := databases.Sqlite()
	jsonResp, _ := json.Marshal(Sqlite.Health())
	_, _ = w.Write(jsonResp)
}

func (a *ArticlesController) Sample(w http.ResponseWriter, r *http.Request) {
	model := models.Article()
	if article, err := model.GetByID(5); err != nil {
		resp := make(map[string]string)
		resp["message"] = err.Error()
		marsh, _ := json.Marshal(resp)
		_, _ = w.Write(marsh)
	} else {
		marsh, _ := json.Marshal(article)
		_, _ = w.Write(marsh)
	}
}
