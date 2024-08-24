package controllers

import (
	"encoding/json"
	"net/http"
	"rixlog/internal/models"

	"github.com/go-chi/chi/v5"
)

func Home() *HomeController {
	if _Home != nil {
		return _Home
	}

	_Home = &HomeController{}
	return _Home
}

type HomeController struct{}

var _Home *HomeController

func (f HomeController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", f.Sample)
	return r
}

// TODO: Create a templates folder with one template for this handler
func (v *HomeController) Sample(w http.ResponseWriter, r *http.Request) {
	if article, err := models.User().GetByID(1); err != nil {
		resp := make(map[string]string)
		resp["message"] = err.Error()
		marsh, _ := json.Marshal(resp)
		_, _ = w.Write(marsh)
	} else {
		marsh, _ := json.Marshal(article)
		_, _ = w.Write(marsh)
	}
}
