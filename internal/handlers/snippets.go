package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/michaelmaysonet74/snpt/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Snippets struct {
	l *log.Logger
	c *mongo.Collection
}

func (s *Snippets) extractID(r *http.Request) (int, error) {
	reg := regexp.MustCompile(`/([0-9]+)$`)

	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) == 0 {
		s.l.Println("ID not found")
		return 0, fmt.Errorf("Invalid URI")
	}

	id, err := strconv.Atoi(g[0][1])
	if err != nil {
		s.l.Println("Invalid URI, unable to convert ID to number")
		return id, fmt.Errorf("Invalid URI")
	}

	return id, nil
}

func (s *Snippets) getSnippets(w http.ResponseWriter, r *http.Request) {
	s.l.Println("GET", r.URL.Path)

	snippets := models.GetSnippets()

	err := models.ToJSON(snippets, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) getSnippetByID(id int, w http.ResponseWriter, r *http.Request) {
	s.l.Println("GET", r.URL.Path)

	snippet, err := models.GetSnippetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := models.ToJSON(snippet, w); err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) createSnippet(w http.ResponseWriter, r *http.Request) {
	s.l.Println("POST", r.URL.Path)
	snippet := &models.Snippet{}

	err := models.FromJSON(snippet, r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	models.CreateSnippet(snippet)
}

func (s *Snippets) updateSnippet(id int, w http.ResponseWriter, r *http.Request) {
	s.l.Println("PUT", r.URL.Path)
	snippet := &models.Snippet{}

	err := models.FromJSON(snippet, r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	models.UpdateSnippet(id, snippet)
}

func (s *Snippets) deleteSnippet(id int, w http.ResponseWriter, r *http.Request) {
	s.l.Println("DELETE", r.URL.Path)

	snippets, err := models.DeleteSnippet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = models.ToJSON(snippets, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusBadRequest)
		return
	}
}

/**
 *	Public Methods
 */

func NewSnippets(l *log.Logger, db *mongo.Database) *Snippets {
	collection := db.Collection("snippets")
	return &Snippets{l, collection}
}

func (s *Snippets) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		s.createSnippet(w, r)

	case http.MethodPut:
		id, err := s.extractID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		s.updateSnippet(id, w, r)

	case http.MethodDelete:
		id, err := s.extractID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.deleteSnippet(id, w, r)

	case http.MethodGet:
		id, err := s.extractID(r)
		if err != nil && id > 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if id > 0 {
			s.getSnippetByID(id, w, r)
		} else {
			s.getSnippets(w, r)
		}

	default:
		s.l.Printf("Method %v is not allowed.", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
