package db

import (
	"context"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

/* ReadTweets reads the tweets from one profile */
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	var response []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	/* Working with options find mode
	limit in 20
	sort is order by date desc*/
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	/* skip 0 tweets in page 1, skip the first 20 tweets for page 2 ... */
	options.SetSkip((page -1)*20)

	/* cursor is like a table to view one by one and processing it */
	cursor, err := collection.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return response, false
	}

	/* T O D O creates a new empty context */
	for cursor.Next(context.TODO()) {
		var register models.ReturnTweets
		err := cursor.Decode(&register)
		if err != nil {
			return response, false
		}
		response = append(response, &register)
	}
	return response, true
}
