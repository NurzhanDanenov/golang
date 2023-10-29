package main

import (
	"log"

	"restapi/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(meet.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())

	}
}
