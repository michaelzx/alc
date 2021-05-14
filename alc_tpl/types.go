package pld_tpl

import "html/template"

type DataMap map[string]interface{}

func NewDataMap() DataMap {
	return make(DataMap)
}

type FuncMap template.FuncMap

func NewFuncMap() FuncMap {
	return make(FuncMap)
}
