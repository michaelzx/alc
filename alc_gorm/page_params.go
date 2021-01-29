package alc_gorm

type PageParams struct {
	PageNum  int64 `valid:"required" cn:"页码"`
	PageSize int64 `valid:"required" cn:"页数"`
}

func (p PageParams) GetPageNum() int64 {
	return p.PageNum
}

func (p PageParams) GetPageSize() int64 {
	return p.PageSize
}
