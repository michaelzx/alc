package alc_sql

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/michaelzx/alc/alc_crypto"
	"github.com/michaelzx/alc/alc_reflect"
	"github.com/patrickmn/go-cache"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ParseTpl(tplStr string, params interface{}) (sqlStr string, sqlParams []interface{}, err error) {
	// ****************************************************
	// 检查params是否是指针
	// ****************************************************
	if !alc_reflect.IsPtr(params) {
		err = NewErr("params 必须是指针类型")
		return
	}
	// ****************************************************
	// 获取模板，避免重复创建性能开销
	// ****************************************************
	tplMd5, err := getTpl(tplStr)
	if err != nil {
		err = NewErr("获取模板失败" + err.Error())
		return
	}
	// ****************************************************
	// 如果cache存在，则直接返回
	// ****************************************************
	cacheKey := tplMd5 + "-" + alc_crypto.Md5(fmt.Sprintf("%#v", params))
	tplCache, found := tplCaches.Get(cacheKey)
	if found {
		if c, exists := tplCache.(TplCache); exists {
			return c.SqlStr, c.SqlParams, nil
		}
	}

	// ****************************************************
	// 解析模板
	// ****************************************************
	var sql bytes.Buffer
	if err = tmpl.ExecuteTemplate(&sql, tplMd5, params); err != nil {
		err = NewErr("解析模板失败" + err.Error())
		return
	}
	sqlStr = sql.String()

	// ****************************************************
	// 对模板中的sql占位符进行解析
	// :xxxx_xxx[x]
	// ****************************************************
	tplParamsRegexp := regexp.MustCompile(`@(?P<param>(?P<name>\w*)(\[(?P<idx>\d+)])?)`)
	tplParamGroups := tplParamsRegexp.FindAllStringSubmatch(sqlStr, -1)
	tplParams := make([]TplParam, 0)
	for _, tplParamGroup := range tplParamGroups {
		tplParams = append(tplParams, TplParam{
			Full:  tplParamGroup[0],
			Name:  tplParamGroup[2],
			Index: tplParamGroup[4],
		})
	}

	// ****************************************************
	// 替换为sql占位符，并按照顺序生成sql参数
	// :xxxx_xxx[x]
	// ****************************************************
	sqlParams = make([]interface{}, 0, len(tplParams))
	if p, ok := params.(*map[string]interface{}); ok {
		pMap := *p
		for _, tplParam := range tplParams {
			v, exist := pMap[tplParam.Name]
			if !exist {
				err = NewErr("params中不存在" + tplParam.Full)
				return
			}
			sqlStr = strings.Replace(sqlStr, tplParam.Full, "?", 1)
			rv := reflect.ValueOf(v)
			if tplParam.Index != "" && (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) {
				idx, _ := strconv.Atoi(tplParam.Index)
				if idx > rv.Len()-1 {
					err = NewErr(tplParam.Full + "超出最大长度")
					return
				}
				paramValue := rv.Index(idx)
				switch paramValue.Kind() {
				case reflect.String:
					sqlParams = append(sqlParams, paramValue.String())
				case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
					sqlParams = append(sqlParams, paramValue.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
					sqlParams = append(sqlParams, paramValue.Uint())
				case reflect.Bool:
					sqlParams = append(sqlParams, paramValue.Bool())
				default:
					err = NewErr("@xxxx[x]不支持:" + paramValue.Kind().String())
					return
				}
			} else {
				sqlParams = append(sqlParams, v)
			}
		}
	} else if alc_reflect.IsStruct(params) {
		rt := reflect.TypeOf(params)
		rve := reflect.ValueOf(params).Elem()
		fim := alc_reflect.StructFieldIdxMap(rt)
		for _, tplParam := range tplParams {
			fi, exist := fim[tplParam.Name]
			if !exist {
				err = NewErr("params中不存在" + tplParam.Full)
				return
			}
			sqlStr = strings.Replace(sqlStr, tplParam.Full, "?", 1)
			v := rve.Field(fi).Interface()
			rv := reflect.ValueOf(v)
			if tplParam.Index != "" && (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) {
				idx, _ := strconv.Atoi(tplParam.Index)
				if idx > rv.Len()-1 {
					err = NewErr(tplParam.Full + "超出最大长度")
					return
				}
				paramValue := rv.Index(idx)
				switch paramValue.Kind() {
				case reflect.String:
					sqlParams = append(sqlParams, paramValue.String())
				case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
					sqlParams = append(sqlParams, paramValue.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
					sqlParams = append(sqlParams, paramValue.Uint())
				case reflect.Bool:
					sqlParams = append(sqlParams, paramValue.Bool())
				default:
					err = NewErr("@xxxx[x]不支持:" + paramValue.Kind().String())
					return
				}
			} else {
				sqlParams = append(sqlParams, v)
			}
		}
	} else {
		err = NewErr("params 仅支持 map 及 struct")
		return
	}
	// ****************************************************
	// 去掉换行、大于1个的空格
	// ****************************************************
	sqlStr = strings.ReplaceAll(sqlStr, "\n", " ")
	re, _ := regexp.Compile(`\s{2,}`)
	sqlStr = re.ReplaceAllString(sqlStr, " ")
	tplCaches.Set(cacheKey, TplCache{
		SqlStr:    sqlStr,
		SqlParams: sqlParams,
	}, cache.NoExpiration)
	return
}

var tmpl = template.New("")

func getTpl(sqlTplStr string) (string, error) {
	tplMd5 := alc_crypto.Md5(sqlTplStr)
	t := tmpl.Lookup(tplMd5)
	if t != nil {
		return tplMd5, nil
	}
	_, err := tmpl.New(tplMd5).Parse(sqlTplStr)
	if err != nil {
		return "", err
	}
	return tplMd5, nil
}
