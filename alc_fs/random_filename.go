package alc_fs

import (
	"alchemy/alc/alc_random"
	"strconv"
	"time"
)

func RandomFilename() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000, 10) + "_" + alc_random.RandomNumStr(6)
}
