package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rixlog/internal/controllers"
	"rixlog/internal/webserver"

	"github.com/go-chi/chi/v5"
)

func Sample() *SampleController {
	if _Sample != nil {
		return _Sample
	}

	_Sample = &SampleController{}
	return _Sample
}

type SampleController struct{}

var _Sample *SampleController

// TODO: redirect to /home
func (s *SampleController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp := make(map[string]string)
		resp["message"] = "Sample Message"
		resp["application"] = "gomvc Sample project"
		marsh, _ := json.Marshal(resp)
		_, _ = w.Write(marsh)
	})
	return r
}

var WebServer = webserver.New(controllers.ControllerMap{
	"/":         Sample(),
	"/home":     controllers.Home(),
	"/articles": controllers.Articles(),
	"/auth":     controllers.Auth(),
	"/admin":    &controllers.Admin{},
})

// TODO: Welsome msg + server info message
func main() {
	fmt.Println("Starting webserver http://localhost:3333")
	if err := WebServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %v", err))
	}
}
