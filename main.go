package main

import (
	"net/http"

	"github.com/enirox001/go-quest/controllers"
	"github.com/enirox001/go-quest/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr: "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
