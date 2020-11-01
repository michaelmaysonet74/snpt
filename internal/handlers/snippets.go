package handlers

import (
	"log"
	"net/http"

	"github.com/michaelmaysonet74/snpt/internal/models"
)

type Snippets struct {
	l *log.Logger
}

func NewSnippets(l *log.Logger) *Snippets {
	return &Snippets{l}
}

func (s *Snippets) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		s.getSnippets(w, r)
		return
	}

	if r.Method == http.MethodPost {
		s.createSnippet(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Snippets) getSnippets(w http.ResponseWriter, r *http.Request) {
	s.l.Println("GET /snippets")
	snippets := models.GetSnippets()

	err := snippets.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) createSnippet(w http.ResponseWriter, r *http.Request) {
	s.l.Println("POST /snippets")
	snippet := &models.Snippet{}

	err := snippet.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	models.CreateSnippet(snippet)
}
