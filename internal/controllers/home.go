package controllers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"rixlog/internal/views"
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
	r.Get("/", f.Index)
	return r
}

func (v *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	views.Index(w, r)
}
