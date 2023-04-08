package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlosarraes/unified/server/controller"
	"github.com/carlosarraes/unified/server/model"
	"github.com/joho/godotenv"
)

func main() {
	port := 8080
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := controller.App{}
	flag.StringVar(&app.DSN, "dsn", os.Getenv("DB_URI"), "PlanetScale DSN")
	flag.Parse()

	conn, err := app.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app.DB = &model.MySql{DB: conn}

	server := app.Routes()

	log.Printf("Server listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
