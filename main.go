package main

import (
	"fmt"
	"log"

	"github.com/FarStep131/go-simple-server/server"
)

const (
	port = 8080
	host = "0.0.0.0"
)

func main() {
	s := server.New()

	err := s.Start(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		s.Stop()
		log.Fatalf("Failed to start server: %v", err)
	}
}
