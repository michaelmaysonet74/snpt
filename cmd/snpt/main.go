package main

import (
	"github.com/michaelmaysonet74/snpt/internal/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
