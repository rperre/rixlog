package webserver

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"rixlog/internal/controllers"
)

func New(routes controllers.RouteMap) *http.Server {
	config := Config{
		Controllers:  routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Port:         3333,
		HTTPS:        false, // TODO: implement
		URL:          "",    // TODO: implement
	}
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      config.RouterHandler(),
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
}

type Config struct {
	Port         int
	URL          string
	HTTPS        bool
	Controllers  controllers.RouteMap
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Config) RouterHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	for name, controller := range s.Controllers {
		r.Mount(name, controller.Routes())
	}
	return r
}
