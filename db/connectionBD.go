package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/* MongoCN is the connection object to the database */
var MongoCN = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://zsfrancisco:3184885238javier@twittercluster.siqz1.mongodb.net/twittor?retryWrites=true&w=majority")

/* ConnectBD is the function to permit the database connection */
func ConnectBD() *mongo.Client  {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("successful connection with db")
	return client
}

/* CheckConnection is the ping to the database */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}