package server

import (
	"blog/internal/storage"

	"html/template"
	"net/http"
	"gopkg.in/yaml.v3"
)

func (s *Server) postFild(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, nil)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	mblog, err := yaml.Marshal(blog)
	if err != nil {
		s.log.Info("Error in marshal blog: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.queue.Add("queue:blogs", mblog)
}