package db

import (
	"context"
	"fmt"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/* ConsultRelation consults the relation between 2 users */
func ConsultRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	collection := db.Collection("relation")

	condition := bson.M{
		"userid": relation.UserID,
		"userrelationid": relation.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)

	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
