package server

import (
	"blog/internal/storage"

	"html/template"
	"net/http"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, nil)
}

func (s *Server) posts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		blog := storage.Blog{
			Author: r.FormValue("author"),
			Title: r.FormValue("title"),
			Body: r.FormValue("body"),
		}
	
		if err := s.storage.Add(blog); err != nil {
			s.log.Info("Error to add in db: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		blogs, err := s.storage.List()
		if err != nil {
			s.log.Info("Error to list in db: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.ParseFiles("./templates/tiles.html"))
		tmpl.Execute(w, blogs)
	} else if r.Method == http.MethodGet {
		blogs, err := s.storage.List()
		if err != nil {
			s.log.Info("Error to list in db: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.ParseFiles("./templates/tiles.html"))
		tmpl.Execute(w, blogs)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (s *Server) getPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	author := r.FormValue("author")

	blogs, err := s.storage.GetPost(title, author)
	if err != nil {
		s.log.Info("Error to list in db: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(blogs) != 0 {
		s.storage.IncreaseViews(title, author)
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("./templates/tiles.html"))
	tmpl.Execute(w, blogs)
}