package alc_sql

import (
	"github.com/patrickmn/go-cache"
)

var tplCaches = cache.New(cache.NoExpiration, cache.NoExpiration)

type TplCache struct {
	SqlStr    string
	SqlParams []interface{}
}
