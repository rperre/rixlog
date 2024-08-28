package views

import (
	"net/http"
	"rixlog/internal/models"
	"text/template"
)

func Article(w http.ResponseWriter, r *http.Request, article *models.Article) {
	tmpl := template.Must(template.ParseFiles("internal/views/article.html"))
	_ = tmpl.Execute(w, article)
}
