package db

import (
	"context"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* InsertTweet saves the tweet in the database */
func InsertTweet(tweet models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	register := bson.M{
		"userid" : tweet.UserID,
		"message": tweet.Message,
		"date": tweet.Date,
	}

	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	/* objID is a tweet id := returns the ultimate register inserted  */
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
