package main

import (
	"Todo/internal/handler"
	"Todo/server"
	"log"
)

func main() {
	handler := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run(":8080", handler.InitRoutes()); err != nil {
		log.Fatalf("err occured while running http server: %s", err.Error())
	}

}
