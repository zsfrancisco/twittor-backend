package db

import (
	"context"
	"github.com/zsfrancisco/twittor-backend/models"
	"time"
)

/* InsertRelation saves the relation in the database */
func InsertRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	collection := db.Collection("relation")

	_, err := collection.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}

	return true, err
}
