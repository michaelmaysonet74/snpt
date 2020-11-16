package server

import (
	"log"
	"os"

	"github.com/michaelmaysonet74/snpt/internal/handlers"
)

func (s *Server) routes() {
	// Dependencies
	l := log.New(os.Stdout, "snpt-api", log.LstdFlags)
	db := NewDB(s.config.DBClientURI, s.config.DBName)

	// Handlers
	sh := handlers.NewSnippets(l, db)

	// Routes
	s.router.Handle("/snippets", sh)
	s.router.Handle("/snippets/", sh)
}
