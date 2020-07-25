package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zsfrancisco/twittor-backend/middlew"
	"github.com/zsfrancisco/twittor-backend/routers"
	"log"
	"net/http"
	"os"
)

/* Handlers the routers of the application */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}