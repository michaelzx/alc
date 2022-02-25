package alc_mgo

import (
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// CNTZ 转成中国时区字符串
func CNTZ(fieldName string) bson.M {
	return bson.M{
		operator.DateToString: bson.M{
			"date":     "$" + fieldName,
			"format":   "%Y-%m-%d %H:%M:%S",
			"timezone": "+08",
		},
	}
}
