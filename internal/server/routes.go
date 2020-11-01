package server

import (
	"log"
	"os"

	"github.com/michaelmaysonet74/snpt/internal/handlers"
)

func (s *Server) routes() {
	// Dependencies
	l := log.New(os.Stdout, "snpt-api", log.LstdFlags)

	// Handlers
	sh := handlers.NewSnippets(l)

	// Routes
	s.router.Handle("/snippets", sh)
}
