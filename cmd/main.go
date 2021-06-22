package main

import (
	"github.com/goodfather13/good-cloud-api-agent"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(agent.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}