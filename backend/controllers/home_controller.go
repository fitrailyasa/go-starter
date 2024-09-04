package controllers

import (
	"html/template"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/client/app.html",
		"frontend/templates/layouts/client/header.html",
		"frontend/templates/layouts/client/footer.html",
		"frontend/templates/pages/client/index.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func AboutController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/client/app.html",
		"frontend/templates/layouts/client/header.html",
		"frontend/templates/layouts/client/footer.html",
		"frontend/templates/pages/client/about.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func ContactController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/client/app.html",
		"frontend/templates/layouts/client/header.html",
		"frontend/templates/layouts/client/footer.html",
		"frontend/templates/pages/client/contact.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}
