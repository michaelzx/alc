package alc_mgo

import (
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
)

const DeletedAtKey = "deleted_at"
const UpdatedAtKey = "updated_at"

var DeletedAtExists = bson.M{
	operator.Exists: true,
}

var DeletedAtNotExists = bson.M{
	operator.Exists: false,
}
