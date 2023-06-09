package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// Home page
	r.Get("/", homeHandler)

	// About page
	r.Get("/about", aboutHandler)

	// Blog post page
	r.Get("/blog/{id}", blogPostHandler)

	var port = envPortOr("3000")
	http.ListenAndServe(port, r)
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and pass any necessary data
	// to be rendered in the template
	data := struct {
		Title string
		// Other data...
	}{
		Title: "Home Page",
		// Other data...
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and pass any necessary data
	// to be rendered in the template
	data := struct {
		Title string
		// Other data...
	}{
		Title: "About Page",
		// Other data...
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func blogPostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and pass any necessary data
	// to be rendered in the template
	data := struct {
		Title string
		// Other data...
	}{
		Title: "Post Page",
		// Other data...
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
