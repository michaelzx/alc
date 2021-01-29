package db_meta

import (
	"database/sql"
	"strings"
)

// Constants for return types of golang
const (
	goString        = "string"
	goDate          = "alc_types.Date"
	goTime          = "alc_types.Time"
	goInt8          = "int8"
	goInt16         = "int16"
	goInt32         = "int32"
	goInt64         = "int64"
	goDecimal       = "decimal.Decimal"
	goBool          = "alc_types.Bool"
	goJson          = "alc_types.JsonString"
	goByteArray     = "[]byte"
	goNullString    = "*string"
	goNullDate      = "*alc_types.Date"
	goNullTime      = "*alc_types.Time"
	goNullInt8      = "*int8"
	goNullInt16     = "*int16"
	goNullInt32     = "*int32"
	goNullInt64     = "*int64"
	goNullDecimal   = "*decimal.Decimal"
	goNullBool      = "*alc_types.Bool"
	goNullJson      = "*alc_types.JsonString"
	goNullByteArray = "*[]byte"
	// todo 确认下是否必须使用NullBool
	// goNullBool      = "alc_types.NullBool"
)

func SqlType2GoType(c *sql.ColumnType) string {
	mysqlType := strings.ToLower(c.DatabaseTypeName())
	nullable, _ := c.Nullable()
	switch mysqlType {
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		if nullable {
			return goNullString
		}
		return goString

	case "tinyint":
		if nullable {
			return goNullInt8
		}
		return goInt8
	case "smallint":
		if nullable {
			return goNullInt16
		}
		return goInt16
	case "mediumint", "int":
		if nullable {
			return goNullInt32
		}
		return goInt32
	case "bigint":
		if nullable {
			return goNullInt64
		}
		return goInt64
	case "decimal", "float", "double":
		if nullable {
			return goNullDecimal
		}
		return goDecimal
	case "datetime", "time", "timestamp":
		if nullable {
			return goNullTime
		}
		return goTime
	case "date":
		if nullable {
			return goNullDate
		}
		return goDate
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		if nullable {
			return goNullByteArray
		}
		return goByteArray
	case "bit":
		if nullable {
			return goNullBool
		}
		return goBool
	case "json":
		if nullable {
			return goNullJson
		}
		return goJson
	}

	return ""
}
