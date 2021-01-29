package example

import "alchemy/alc/alc_errs"

var (
	DataNotFound   = alc_errs.NewBadRequest(10001, "数据不存在")
	DataDuplicated = alc_errs.NewBadRequest(10002, "已被占用")
	// ...
)
