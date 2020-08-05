package db

import (
	"context"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* ModifyRegister allows modify the user profile */
func ModifyRegister(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("users")

	register := make(map[string]interface{})

	// Checking values modified
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Surname) > 0 {
		register["surname"] = user.Surname
	}
	register["birthday"] = user.Birthday
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}
	if len(user.Website) > 0 {
		register["website"] = user.Website
	}

	/* Building the json to update the mongo record
	In mongo, is used a $set to modify one record */
	updateStrig := bson.M{
		"$set" : register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	/* Getting the user in db, $eq is equal than */
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filter, updateStrig)
	if err != nil {
		return false, err
	}

	return true, nil

}
