package alc_snowflake

import (
	"alchemy/alc/alc_errs"
	"github.com/bwmarrin/snowflake"
)

func NewNode(number int64) (node *snowflake.Node, err error) {
	node, err = snowflake.NewNode(number)
	if err != nil {
		err = alc_errs.Wrap(err, "snowflake节点创建失败")
		return
	}
	return
}
