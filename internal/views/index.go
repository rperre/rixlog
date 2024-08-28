package views

import (
	"net/http"
	"text/template"
)

type IndexView struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/views/index.html"))
	data := &IndexView{Title: "We made it"}
	_ = tmpl.Execute(w, data)
}
