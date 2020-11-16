package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server holds everything needed by the API
type Server struct {
	config *Config
	router *http.ServeMux
}

/**
 *	Public Methods
 */

func NewServer() *Server {
	return &Server{
		config: NewConfig(),
		router: http.NewServeMux(),
	}
}

// Run is used as the service initializer
func (s *Server) Run() {
	s.routes()

	server := &http.Server{
		Addr:         s.config.Addr,
		Handler:      s.router,
		IdleTimeout:  s.config.IdleTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}

	log.Printf(
		"Starting server on port %s",
		s.config.Addr,
	)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
