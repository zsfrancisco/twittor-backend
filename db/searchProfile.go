package db

import (
	"context"
	"fmt"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* SearchProfile searches a profile in the db */
func SearchProfile(ID string) (models.User, error)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		fmt.Println("Register not found "+err.Error())
		return profile, err
	}

	return profile, nil

}
