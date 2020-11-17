package server

import (
	"log"
	"os"

	"github.com/michaelmaysonet74/snpt/internal/handlers"
)

func (s *Server) routes() {
	// Dependencies
	l := log.New(os.Stdout, "snpt-api", log.LstdFlags)
	db := s.db

	// Handlers
	sh := handlers.NewSnippets(l, db)

	// Routes
	s.router.HandleFunc("/snippets", sh.GetSnippets).Methods("GET")
	s.router.HandleFunc("/snippets", sh.CreateSnippet).Methods("POST")

	s.router.HandleFunc("/snippets/{id}", sh.GetSnippetByID).Methods("GET")
	s.router.HandleFunc("/snippets/{id}", sh.UpdateSnippet).Methods("PUT")
	s.router.HandleFunc("/snippets/{id}", sh.DeleteSnippet).Methods("DELETE")
}
