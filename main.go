package main

import (
	"log"

	"github.com/reubenthomasjohn/location-heatmap/api"
)

func main() {
	server := api.NewServer()

	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}