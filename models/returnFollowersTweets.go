package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* ReturnFollowersTweets is the struct used to return tweets */
type ReturnFollowersTweets struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string `bson:"userid" json:"userId,omitempty"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId,omitempty"`
	Tweet struct{
		Message string `bson:"message" json:"message,omitempty"`
		Date time.Time `bson:"date" json:"date,omitempty"`
		ID string `bson:"_id" json:"_id,omitempty"`
	}
}
