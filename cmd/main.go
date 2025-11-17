package main

import (
	"log"

	"github.com/incheat/go-mastermind/internal/server"
)

func main() {
	r := server.New()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
