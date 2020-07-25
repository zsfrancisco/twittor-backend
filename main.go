package main

import (
	"github.com/zsfrancisco/twittor-backend/db"
	"github.com/zsfrancisco/twittor-backend/handlers"
	"log"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No connection with database")
		return
	}
	handlers.Handlers()
}
