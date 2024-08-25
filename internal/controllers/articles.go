package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rixlog/internal/models"
	"rixlog/internal/views"
	"strconv"
	"strings"

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
	r.With(paginate).Get("/", a.GetArticle)
	r.With(Authenticated).Post("/", a.Void)
	r.With(Authenticated).Put("/", a.Void)
	r.With(Authenticated).With(AdminOnly).Delete("/", a.Void)
	return r
}

func (a *ArticlesController) GetArticle(w http.ResponseWriter, r *http.Request) {
	model := models.Article()
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if article, err := model.GetByID(id); err != nil {
		HtppRerror(err, w, r)
	} else {
		accpet_header := r.Header["Accept"][0]
		if strings.Contains(accpet_header, "text/html") {
			views.Article(w, r, article)
		} else if strings.Contains(accpet_header, "application/json") {
			marsh, _ := json.Marshal(article)
			_, _ = w.Write(marsh)
		}
	}
}

func (a *ArticlesController) Void(w http.ResponseWriter, r *http.Request) {}

func HtppRerror(err error, w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = err.Error()
	marsh, _ := json.Marshal(resp)
	_, _ = w.Write(marsh)
}
