package server

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("ui", "out", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
