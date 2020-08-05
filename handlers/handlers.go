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
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/view_profile", middlew.CheckBD(middlew.ValidJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify_profile", middlew.CheckBD(middlew.ValidJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidJWT(routers.SaveTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}