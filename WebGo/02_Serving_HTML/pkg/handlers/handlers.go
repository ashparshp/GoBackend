package handlers

import (
	"net/http"
	"templates/pkg/render"
)

func HomePage (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
func AboutPage (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}