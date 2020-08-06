package models

/* Relation is the model for save the relation from an user with other */
type Relation struct {
	UserID string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}
