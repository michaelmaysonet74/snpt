package main

import (
	"log"

	"github.com/michaelmaysonet74/snpt/internal/server"
)

func main() {
	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}
