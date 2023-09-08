package main

import (
	"os"

	"github.com/andrerocco/rinha-de-backend-go/database"
	"github.com/andrerocco/rinha-de-backend-go/http_api"
)

func main() {
	store, err := database.Connect()
	if err != nil {
		panic(err)
	}

	server := http_api.NewAPIServer(os.Getenv("API_PORT"), store)
	server.Start()
}
