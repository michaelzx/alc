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
	PageNum     int64 // 第几页
	PageSize    int64 // 每页几条
	PageTotal   int64 // 总共几页
	Total       int64 // 总共几条
	IsFirstPage bool  // 是否是第一页
	IsLastPage  bool  // 是否是最后一页
	List        interface{}
}

func NewPageVO(db *gorm.DB, list interface{}, sqlTpl string, params IPage) (*PageVO, error) {
	// return &PageVO{PageNum: pageNum, PageSize: pageSize, List: list}
	p := &PageVO{
		PageNum:  params.GetPageNum(),
		PageSize: params.GetPageSize(),
		List:     list,
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
	countSql := fmt.Sprintf(`select count(*) from (%s) as t`, sqlStr)
	result := db.Raw(countSql, sqlParams...).Count(&p.Total)
	if result.Error != nil {
		return result.Error
	}
	p.PageTotal = p.Total / p.PageSize
	if p.Total%p.PageSize > 0 {
		p.PageTotal = p.PageTotal + 1
	}
	if p.PageNum > p.PageTotal {
		p.PageNum = p.PageTotal
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if p.PageTotal == 0 {

		p.IsFirstPage = true
		p.IsLastPage = true
	} else {
		if p.PageNum == 1 {
			p.IsFirstPage = true
		}
		if p.PageNum == p.PageTotal {
			p.IsLastPage = true
		}
	}
	skipRow := p.PageSize * (p.PageNum - 1)
	pageSql := fmt.Sprintf(`%s limit %d,%d`, sqlStr, skipRow, p.PageSize)
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
