package alc_mgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"time"
)

var nilTime time.Time

// DefaultField defines the default fields to handle when operation happens
// import the DefaultField in document struct to make it working
type DefaultField struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// DefaultUpdateAt changes the default updateAt field
func (df *DefaultField) DefaultUpdateAt() {
	df.UpdatedAt = time.Now().Local()
}

// DefaultCreateAt changes the default createAt field
func (df *DefaultField) DefaultCreateAt() {
	if reflect.DeepEqual(df.CreatedAt, nilTime) {
		df.CreatedAt = time.Now().Local()
	}
}

// DefaultId changes the default _id field
func (df *DefaultField) DefaultId() {
	if df.ID.IsZero() {
		df.ID = primitive.NewObjectID()
	}
}
