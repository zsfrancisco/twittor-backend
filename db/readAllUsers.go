package db

import (
	"context"
	"fmt"
	"github.com/zsfrancisco/twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/* ReadAllUsers reads the users registered in the system,
if get "R", returns the relations with me only*/
func ReadAllUsers(ID string, page int64, search string, typeSearch string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page-1)*2)
	findOptions.SetLimit(20)

	/* comparing the search string with $regex regular expression (?i) i is that doesn't import if is upper or lower case */
	query := bson.M{
		"name": bson.M{"$regex": `(?i)`+ search},
	}

	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	/* saving the fields in model user */
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relation models.Relation
		relation.UserID = ID
		relation.UserRelationID = user.ID.Hex()

		include = false

		found, err = ConsultRelation(relation)
		/* users that I do not follow */
		if typeSearch == "new" && found == false {
			include = true
		}
		/* users that I follow */
		if typeSearch == "follow" && found == true {
			include = true
		}

		if relation.UserRelationID == ID {
			include = false
		}

		if include == true {
			user.Password = ""
			user.Biography = ""
			user.Website = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	/* if there's an error while going through the cursor */
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	/* Closing the cursor */
	cursor.Close(ctx)
	return results, true
}
