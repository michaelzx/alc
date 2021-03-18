package alc_gen

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"
	"github.com/michaelzx/alc/alc_arrays"
	"github.com/michaelzx/alc/alc_config"
	"github.com/michaelzx/alc/alc_fs"
	"github.com/michaelzx/alc/alc_gen/db_meta"
	"log"
	"os"
	"strings"
)

type NameProcessor func(structName string) string
type TableMeta struct {
	TableName      string
	StructName     string
	Pks            []string
	ColumnComments map[string]string
	Fields         []string
	EntImports     []string
}

type Generator struct {
	DbCfg               alc_config.MysqlConfig
	TypePackage         string
	Tables              []string
	OutDir              string
	PackageEnt          string
	PackageService      string
	TplEnt              string
	TplService          string
	GenService          bool
	GenEnt              bool
	StructNameProcessor NameProcessor // struct名字自定义处理器
	FileNameProcessor   NameProcessor // 文件名自定义处理器
	// 私有属性，运行过程中用到
	db           *sql.DB
	rootPath     string
	tableMetaMap map[string]TableMeta
}

func (g *Generator) Gen() {
	if len(g.Tables) == 0 {
		panic(errors.New("未指定任何需要生成的表"))
	}
	alc_fs.CreateIfNotExist(g.OutDir)
	sqlType := "mysql"
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&parseTime=True",
		g.DbCfg.Usr,
		g.DbCfg.Psw,
		g.DbCfg.Host,
		g.DbCfg.Port,
		g.DbCfg.DbName,
	)
	fmt.Println(connStr)
	// 创建数据库链接
	if newDb, err := sql.Open(sqlType, connStr); err != nil {
		fmt.Println("数据库链接创建失败: " + err.Error())
		return
	} else {
		g.db = newDb
	}
	defer g.db.Close()
	// 获取程序运行的目录
	g.rootPath, _ = os.Getwd()
	fmt.Println(g.rootPath)
	// 检查表是否在数据库中
	g.checkTables()
	// 从数据库中获取表字段信息
	g.loadTableMeta()
	if g.GenEnt {
		g.genEnt()
	}
	if g.GenService {
		g.genService()
	}
}

func (g *Generator) loadTableMeta() {
	g.tableMetaMap = make(map[string]TableMeta)
	for _, tableName := range g.Tables {
		pks := g.getTablePks(tableName)
		comments := g.getTabComments(tableName)
		fields := g.getFields(tableName, pks, comments)
		entImports := g.getEntImports(fields)
		structName := db_meta.StructName(tableName)
		if g.StructNameProcessor != nil {
			structName = g.StructNameProcessor(structName)
		}
		g.tableMetaMap[tableName] = TableMeta{
			TableName:      tableName,
			StructName:     structName,
			Pks:            pks,
			ColumnComments: comments,
			Fields:         fields,
			EntImports:     entImports,
		}
	}
}

func (g *Generator) getEntImports(fields []string) []string {
	// fields := generateMysqlTypes(db, columnTypes, 0, jsonAnnotation, gormAnnotation, gureguTypes)
	hasAppTypes := false
	for _, field := range fields {
		if strings.Contains(field, "types.") && hasAppTypes == false {
			hasAppTypes = true
			break
		}
	}
	hasDecimal := false
	for _, field := range fields {
		if strings.Contains(field, "decimal.") && hasDecimal == false {
			hasDecimal = true
			break
		}
	}
	hasGorm := false
	for _, field := range fields {
		if strings.Contains(field, "gorm.") && hasDecimal == false {
			hasGorm = true
			break
		}
	}
	var imports []string
	if hasAppTypes {
		imports = append(imports, `"`+g.TypePackage+`"`)
	}
	if hasDecimal {
		imports = append(imports, `"github.com/shopspring/decimal"`)
	}
	if hasGorm {
		imports = append(imports, `"gorm.io/gorm"`)
	}
	return imports
}
func (g *Generator) getFields(tableName string, pks []string, comments map[string]string) []string {
	autoKeys := g.getAutoKeys(tableName)
	var fields []string
	columns, _ := schema.Table(g.db, tableName)
	for _, c := range columns {
		var field = ""
		key := c.Name()
		valueType := db_meta.SqlType2GoType(c)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := db_meta.StructName(db_meta.StringifyFirstChar(key))
		var tagInfoItems []string
		if alc_arrays.IndexOfString(pks, key) >= 0 {
			tagInfoItems = append(tagInfoItems, "pk")
		}
		if alc_arrays.IndexOfString(autoKeys, key) >= 0 {
			tagInfoItems = append(tagInfoItems, "auto")
		}
		if key == "deleted_at" {
			valueType = "gorm.DeletedAt"
		}
		var annotations []string
		if len(tagInfoItems) == 0 {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s\"", key))
		} else {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s;%s\"", key, strings.Join(tagInfoItems, ";")))
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			field = fmt.Sprintf("%s %s",
				fieldName,
				valueType)
		}
		comment, _ := comments[key]
		field += " // " + comment
		fields = append(fields, field)
	}
	return fields
}

func (g *Generator) checkTables() {
	ts, err := schema.TableNames(g.db)
	if err != nil {
		fmt.Println("Error in fetching tables information from mysql information schema")
		return
	}
	for _, table := range g.Tables {
		if alc_arrays.IndexOfString(ts, table) == -1 {
			log.Fatal("数据库中不存在表：" + table)
		}
	}
}

func (g *Generator) getTabComments(tableName string) map[string]string {
	sql := `select column_name,column_comment from INFORMATION_SCHEMA.Columns where table_schema=? and table_name =?`
	rows, err := g.db.Query(sql, g.DbCfg.DbName, tableName)
	if err != nil {
		return nil
	}
	commentMap := make(map[string]string)
	for rows.Next() {
		columnName := ""
		columnComment := ""
		err = rows.Scan(
			&columnName,
			&columnComment,
		)
		if err != nil {
			return nil
		}
		commentMap[columnName] = columnComment
	}
	return commentMap
}

func (g *Generator) getAutoKeys(tableName string) []string {
	rows, err := g.db.Query("describe " + tableName)
	if err != nil {
		panic(err)
	}
	autoKeys := make([]string, 0, 0)
	for rows.Next() {
		var Field, Type, Null, Key, Default, Extra sql.NullString
		rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
		if Extra.String == "auto_increment" {
			autoKeys = append(autoKeys, Field.String)
		}
	}
	return autoKeys
}
func (g *Generator) getTablePks(tableName string) []string {
	var primaryKeysQuery = `select k.column_name
from information_schema.table_constraints t
left join information_schema.key_column_usage k
using(constraint_name,table_schema,table_name)
where t.constraint_type='primary key'
	and t.table_schema = ? and k.table_name=?;
`
	rows, err := g.db.Query(primaryKeysQuery, g.DbCfg.DbName, tableName)
	if err != nil {
		return nil
	}
	pks := make([]string, 0, 0)
	for rows.Next() {
		var pkName string
		err = rows.Scan(
			&pkName,
		)
		if err != nil {
			return nil
		}
		pks = append(pks, pkName)
	}
	return pks
}
