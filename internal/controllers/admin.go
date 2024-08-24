package controllers

import (
	"encoding/json"
	"net/http"
	"rixlog/internal/models"

	"github.com/go-chi/chi/v5"
)

type Admin struct{}

func (a *Admin) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(Authenticated)
	r.Use(AdminOnly)
	r.Get("/db_health", a.Sample)
	return r
}

func (v *Admin) Sample(w http.ResponseWriter, r *http.Request) {
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
