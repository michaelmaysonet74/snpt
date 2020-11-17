package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

// Server holds everything needed by the API
type Server struct {
	config *Config
	router *mux.Router
	db     *mongo.Database
}

/**
 *	Public Methods
 */

func NewServer() *Server {
	return &Server{
		config: NewConfig(),
		router: mux.NewRouter(),
	}
}

// Run is used as the service initializer
func (s *Server) Run() {
	client, ctx := s.NewDBClient()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	s.db = client.Database(s.config.DBName)
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
