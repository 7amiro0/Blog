package server

import (
	"blog/internal/storage"
	"time"

	"html/template"
	"net/http"
)

const (
	htmlIndex = "./templates/index.html"
	lifetimeCache = time.Second * 60
)

func (s *Server) createPost(w http.ResponseWriter, r *http.Request) {
	blog := storage.Blog{
		Author: r.FormValue("author"),
		Title:  r.FormValue("title"),
		Body:   r.FormValue("body"),
	}

	if err := s.storage.Add(blog); err != nil {
		s.log.Info("Error to add in db: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) listPosts(w http.ResponseWriter, r *http.Request) {
	blogs, err := s.cache.Get("blogs")
	if err == nil {
		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.ParseFiles(htmlIndex))
		tmpl.Execute(w, blogs.GetPosts())

		return
	}
	
	blogs, err = s.storage.List()
	if err != nil {
		s.log.Info("Error to list in db: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.cache.Set("blogs", blogs, time.Duration(lifetimeCache))
	if err != nil {
		s.log.Info("Error cant save cache: ", err)
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles(htmlIndex))
	tmpl.Execute(w, blogs.GetPosts())
}

func (s *Server) getPost(w http.ResponseWriter, r *http.Request, title, author string) {
	blogs, err := s.cache.Get(title+author)
	if err == nil {
		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.ParseFiles(htmlIndex))
		tmpl.Execute(w, blogs.GetPosts())

		return
	}
	
	blogs, err = s.storage.GetPost(title, author)
	if err != nil {
		s.log.Info("Error to list in db: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.cache.Set(title+author, blogs, time.Duration(lifetimeCache))
	if err != nil {
		s.log.Info("Error cant save cache: ", err)
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles(htmlIndex))
	tmpl.Execute(w, blogs.GetPosts())
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.createPost(w, r)
		s.listPosts(w, r)
	} else if r.Method == http.MethodGet {
		title := r.FormValue("title")
		author := r.FormValue("author")
		
		if title != "" && author != "" {
			s.getPost(w, r, title, author)
		} else {
			s.listPosts(w, r)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}