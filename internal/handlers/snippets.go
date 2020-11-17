package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/michaelmaysonet74/snpt/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Snippets struct {
	l *log.Logger
	c *mongo.Collection
}

func (s *Snippets) extractID(r *http.Request) primitive.ObjectID {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	return id
}

/**
 *	Public Methods
 */

func NewSnippets(l *log.Logger, db *mongo.Database) *Snippets {
	collection := db.Collection("snippets")
	return &Snippets{l, collection}
}

func (s *Snippets) GetSnippets(w http.ResponseWriter, r *http.Request) {
	s.l.Println("GET", r.URL.Path)
	w.Header().Add("Content-Type", "application/json")

	snippets, err := models.GetSnippets(s.c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = models.ToJSON(snippets, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) GetSnippetByID(w http.ResponseWriter, r *http.Request) {
	s.l.Println("GET", r.URL.Path)
	w.Header().Add("Content-Type", "application/json")

	id := s.extractID(r)
	snippet, err := models.GetSnippetByID(s.c, id)
	if err != nil {
		http.Error(
			w,
			"Snippet with ID["+id.Hex()+"] not found",
			http.StatusNotFound,
		)
		return
	}

	if err := models.ToJSON(snippet, w); err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	s.l.Println("POST", r.URL.Path)
	w.Header().Add("Content-Type", "application/json")

	snippet := &models.Snippet{}
	err := models.FromJSON(snippet, r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	snippet, err = models.CreateSnippet(s.c, snippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = models.ToJSON(snippet, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) UpdateSnippet(w http.ResponseWriter, r *http.Request) {
	s.l.Println("PUT", r.URL.Path)
	w.Header().Add("Content-Type", "application/json")

	id := s.extractID(r)
	snippet := &models.Snippet{}

	err := models.FromJSON(snippet, r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	result, err := models.UpdateSnippet(s.c, id, snippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = models.ToJSON(result, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (s *Snippets) DeleteSnippet(w http.ResponseWriter, r *http.Request) {
	s.l.Println("DELETE", r.URL.Path)
	w.Header().Add("Content-Type", "application/json")

	id := s.extractID(r)
	err := models.DeleteSnippet(s.c, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
