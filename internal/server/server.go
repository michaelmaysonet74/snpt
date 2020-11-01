package server

import (
	"log"
	"net/http"
	"os"

	"github.com/michaelmaysonet74/snpt/internal/handlers"
)

// Server holds everything needed by the API
type Server struct {
	config *Config
}

func NewServer() *Server {
	return &Server{
		config: NewConfig(),
	}
}

// Run is used as the service initializer
func (s *Server) Run() (e error) {
	l := log.New(os.Stdout, "snpt-api", log.LstdFlags)

	eh := handlers.NewEcho(l)

	smux := http.NewServeMux()
	smux.Handle("/echo", eh)

	log.Printf("Starting server on port %s", s.config.Port)
	return http.ListenAndServe(s.config.Port, smux)
}
