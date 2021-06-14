package alc_gen

import (
	"database/sql"
	"fmt"
	"github.com/jimsmart/schema"
	"github.com/michaelzx/alc/alc_arrays"
	"strings"
)

type ColumnComments = map[string]string

type TableMeta struct {
	db             *sql.DB
	dbName         string
	TableName      string
	StructName     string
	PrimaryKeys    []string
	AutoKeys       []string
	ColumnComments ColumnComments
	Fields         []string
	Imports        []string
}

func NewTableMeta(db *sql.DB, dbName string, tableName string) (*TableMeta, error) {
	structName := structName(tableName)
	t := &TableMeta{
		TableName:  tableName,
		StructName: structName,
		db:         db,
		dbName:     dbName,
	}
	// 解析自增字段
	err := t.parseAutoKeys()
	if err != nil {
		return nil, err
	}
	// 解析主键
	err = t.parsePKs()
	if err != nil {
		return nil, err
	}
	// 解析备注说明
	err = t.parseComments()
	if err != nil {
		return nil, err
	}
	// 解析字段
	t.parseFields()
	// 解析import
	t.parseImports()
	return t, nil
}

const sqlTablePrimaryKeys = `select k.column_name
from information_schema.table_constraints t
left join information_schema.key_column_usage k
using(constraint_name,table_schema,table_name)
where t.constraint_type='primary key'
	and t.table_schema = ? and k.table_name=?;
`

func (t *TableMeta) parsePKs() error {
	rows, err := t.db.Query(sqlTablePrimaryKeys, t.dbName, t.TableName)
	defer rows.Close()
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
			return err
		}
		pks = append(pks, pkName)
	}
	t.PrimaryKeys = pks
	return nil
}

const sqlTableComments = `select column_name,column_comment from information_schema.columns where table_schema=? and table_name =?`

func (t *TableMeta) parseComments() error {
	rows, err := t.db.Query(sqlTableComments, t.dbName, t.TableName)
	defer rows.Close()
	if err != nil {
		return err
	}
	comments := make(ColumnComments)
	for rows.Next() {
		columnName := ""
		columnComment := ""
		err = rows.Scan(
			&columnName,
			&columnComment,
		)
		if err != nil {
			return err
		}
		comments[columnName] = columnComment
	}
	t.ColumnComments = comments
	return nil
}

func (t *TableMeta) parseAutoKeys() error {
	rows, err := t.db.Query("describe " + t.TableName)
	defer rows.Close()
	if err != nil {
		return err
	}
	autoKeys := make([]string, 0, 0)
	for rows.Next() {
		var Field, Type, Null, Key, Default, Extra sql.NullString
		rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
		if Extra.String == "auto_increment" {
			autoKeys = append(autoKeys, Field.String)
		}
	}
	t.AutoKeys = autoKeys
	return err
}

func (t *TableMeta) parseFields() {
	var fields []string
	columns, _ := schema.Table(t.db, t.TableName)
	for _, c := range columns {
		var field = ""
		key := c.Name()
		valueType := sqlType2GoType(c)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := structName(stringifyFirstChar(key))
		var tagInfoItems []string
		if alc_arrays.IndexOfString(t.PrimaryKeys, key) >= 0 {
			tagInfoItems = append(tagInfoItems, "pk")
		}
		if alc_arrays.IndexOfString(t.AutoKeys, key) >= 0 {
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
		comment, _ := t.ColumnComments[key]
		field += " // " + comment
		fields = append(fields, field)
	}
	t.Fields = fields
}

func (t *TableMeta) parseImports() {
	// fields := generateMysqlTypes(db, columnTypes, 0, jsonAnnotation, gormAnnotation, gureguTypes)
	hasAppTypes := false
	for _, field := range t.Fields {
		if strings.Contains(field, "alc_types.") && hasAppTypes == false {
			hasAppTypes = true
			break
		}
	}
	hasDecimal := false
	for _, field := range t.Fields {
		if strings.Contains(field, "decimal.") && hasDecimal == false {
			hasDecimal = true
			break
		}
	}
	hasGorm := false
	for _, field := range t.Fields {
		if strings.Contains(field, "gorm.") && hasDecimal == false {
			hasGorm = true
			break
		}
	}
	var imports []string
	if hasAppTypes {
		imports = append(imports, `"github.com/michaelzx/alc/alc_types"`)
	}
	if hasDecimal {
		imports = append(imports, `"github.com/shopspring/decimal"`)
	}
	if hasGorm {
		imports = append(imports, `"gorm.io/gorm"`)
	}
	t.Imports = imports
}
