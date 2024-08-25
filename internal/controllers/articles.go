package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"rixlog/internal/databases"
	"rixlog/internal/models"
	"rixlog/internal/views"
	"strconv"
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
	r.With(paginate).Get("/", a.GetArticle)
	r.Get("/{slug}", a.SampleJSON)
	r.With(Authenticated).Post("/", a.SampleJSON)
	r.With(Authenticated).Put("/{slug}", a.SampleJSON)
	r.With(Authenticated).With(AdminOnly).Delete("/{slug}", a.SampleJSON)
	return r
}

func (a *ArticlesController) HandleDBHealth(w http.ResponseWriter, r *http.Request) {
	Sqlite := databases.Sqlite()
	jsonResp, _ := json.Marshal(Sqlite.Health())
	_, _ = w.Write(jsonResp)
}

func (a *ArticlesController) GetArticle(w http.ResponseWriter, r *http.Request) {
	model := models.Article()
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if article, err := model.GetByID(id); err != nil {
		resp := make(map[string]string)
		resp["message"] = err.Error()
		marsh, _ := json.Marshal(resp)
		_, _ = w.Write(marsh)
	} else {
		views.Article(w, r, article)
	}
}
func (a *ArticlesController) SampleJSON(w http.ResponseWriter, r *http.Request) {
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
