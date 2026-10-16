package main

import (
	"hook-flow/internal/bootstrap"
	"log"
)

func main() {

	if err := bootstrap.Start(); err != nil {
		log.Fatal(err)
	}
}
