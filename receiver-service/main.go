package main

import (
	"log"
	"net/http"
	"receiver-service/api"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/vehicle", api.HandleVehicleReceived).Methods("POST")

	log.Println("Server is running...")
	http.ListenAndServe(":8085", r)
}
