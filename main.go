package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").ParseFiles("templates/index.html")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("failed to parse template index: %v", err)))
			return
		}

		tmpl.ExecuteTemplate(w, "Base", nil)
	})
	http.ListenAndServe("localhost:3000", r)
}
