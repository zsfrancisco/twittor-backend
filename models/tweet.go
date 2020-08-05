package models

/* Tweet gets the message from the body of the frontend */
type Tweet struct {
	Message string `bson:"message" json:"message, omitempty"`
}
