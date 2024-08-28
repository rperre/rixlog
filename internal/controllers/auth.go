package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"rixlog/internal/models"
)

func Auth() *AuthController {
	if _Auth != nil {
		return _Auth
	}

	_Auth = &AuthController{}
	return _Auth
}

type AuthController struct{}

var _Auth *AuthController

func (a *AuthController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/login", a.Sample)
	r.Post("/login", a.Sample)
	r.Get("/create_account", a.Sample)
	r.Post("/create_account", a.Sample)
	r.With(Authenticated).Post("/confirm_account", a.Sample)
	r.With(Authenticated).Post("/logout", a.Sample)
	r.With(Authenticated).Get("/change_password", a.Sample)
	r.With(Authenticated).Post("/change_password", a.Sample)
	return r
}

func (v *AuthController) Sample(w http.ResponseWriter, r *http.Request) {
	model := &models.Article{}
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
