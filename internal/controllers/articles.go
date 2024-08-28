package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"rixlog/internal/models"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5"
	"rixlog/internal/views"
	"strconv"
	"strings"
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

func (a *ArticlesController) ExistingArticle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		model := models.Article()
		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		fmt.Println(chi.URLParam(r, "id"))
		fmt.Println(id)
		if article, err := model.GetByID(id); err != nil {
			HttpError(err, http.StatusNotFound, w, r)
			return
		} else {
			ctx := context.WithValue(r.Context(), "article", article)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

type ArticleResponse struct {
	*models.ArticleJSON

	Elapsed int64 `json:"elapsed"`
}

func (rd *ArticleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 0
	return nil
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
			render.Status(r, http.StatusOK)
			_ = render.Render(w, r, &ArticleResponse{ArticleJSON: article})
			_, _ = w.Write(marsh)
		}
	}
}

func (a *ArticlesController) Void(w http.ResponseWriter, r *http.Request) {}

func HttpError(err error, code int, w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = err.Error()
	marsh, _ := json.Marshal(resp)
	_, _ = w.Write(marsh)
}
