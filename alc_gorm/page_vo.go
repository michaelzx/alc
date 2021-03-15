package alc_gorm

import (
	"alchemy/alc/alc_reflect"
	"alchemy/alc/alc_sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type IPage interface {
	GetPageNum() int64
	GetPageSize() int64
}

type PageVO struct {
	Pagination
	List interface{}
}

func NewPageVO(db *gorm.DB, list interface{}, sqlTpl string, params IPage) (*PageVO, error) {
	// return &PageVO{PageNum: pageNum, PageSize: pageSize, List: list}
	p := &PageVO{
		Pagination: NewPagination(params.GetPageNum(), params.GetPageSize()),
		List:       list,
	}
	err := p.get(db, sqlTpl, params)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PageVO) get(db *gorm.DB, sqlTpl string, params interface{}) error {
	// 先用text/template对sql解析一波
	sqlStr, sqlParams, err := alc_sql.ParseTpl(sqlTpl, params)
	if err != nil {
		return err
	}
	// 统计总数
	countSql := fmt.Sprintf(`select count(*) from (%s) as t`, sqlStr)
	var total int64
	result := db.Raw(countSql, sqlParams...).Count(&total)
	if result.Error != nil {
		return result.Error
	}
	p.Pagination.Compute(total)
	pageSql := fmt.Sprintf(`%s limit %d,%d`, sqlStr, p.GetSkipRows(), p.PageSize)
	pageSql = strings.Replace(pageSql, "\n", " ", -1)
	result = db.Raw(pageSql, sqlParams...)
	if !CheckResult(result) {
		return nil
	}
	if !alc_reflect.IsPtr(p.List) {
		return errors.New("字段 List 必须是指针类型")
	}
	result.Scan(p.List)
	return nil
}
