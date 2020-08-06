package db

import (
	"context"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/* ReadFollowersTweets reads the tweets of my followers */
func ReadFollowersTweets(ID string, page int) ([]models.ReturnFollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relation")

	skip := (page - 1) * 20

	/* making []bson with 0 elements */
	conditions := make([]bson.M,0)

	/* Framework aggregate
		$match search the id of the relation -> userid should be equal than ID
		$lookup, allows join 2 tables -> 4 parameters
			from: where i'm joining
			localField field used to join
			foreignField the field in the table tweet in the example -> userid
			as is the alias for the table
	*/

	conditions = append(conditions, bson.M{"$match": bson.M{"userid":ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from": "tweet",
			"localField": "userrelationid",
			"foreignField": "userid",
			"as": "tweet",
		},
	})

	/* unwind used in tweet table
		without unwind:
			marianoRelation { []tweets}
		with unwind:
			all the documents are equals, tweet in the directly document
	*/
	conditions = append(conditions, bson.M{"$unwind":"$tweet"})
	/* order by date desc */
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, conditions)

	var results []models.ReturnFollowersTweets

	/* builds the necessary format that I defined and go through the cursor and assigns it in result decoding that */
	err = cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}
	return results, true
}