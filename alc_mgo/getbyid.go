package alc_mgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromStr(id string) primitive.ObjectID {
	if id == "" {
		return primitive.NilObjectID
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID
	}
	return objectID
}
