package main

import (
	"log"
	"net/http"

	"github.com/HermanPrimo/api_zapateriago/data"
	"github.com/joho/godotenv"

	//"github.com/HermanPrimo/api_golang/models"
	"github.com/HermanPrimo/api_zapateriago/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env file")
	}
	data.ConectarPostgres()
	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":3001", rutas))
}
