package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* DeleteTweet deletes an specific tweet */
func DeleteTweet(ID string, userID string) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	/* converting tweet id */
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
		"userid": userID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}
