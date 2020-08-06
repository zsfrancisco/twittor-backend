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
	router.HandleFunc("/read_tweets", middlew.CheckBD(middlew.ValidJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/delete_tweet", middlew.CheckBD(middlew.ValidJWT(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/upload_avatar", middlew.CheckBD(middlew.ValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/get_avatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/upload_banner", middlew.CheckBD(middlew.ValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/get_banner", middlew.CheckBD(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/high_relation", middlew.CheckBD(middlew.ValidJWT(routers.HighRelation))).Methods("POST")
	router.HandleFunc("/low_relation", middlew.CheckBD(middlew.ValidJWT(routers.LowRelation))).Methods("DELETE")
	router.HandleFunc("/consult_relation", middlew.CheckBD(middlew.ValidJWT(routers.ConsultRelation))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}