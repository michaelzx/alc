package db_meta

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"testing"
)

func TestStructName(t *testing.T) {
	fmt.Println(StructName("_"))
	fmt.Println(StructName("-"))
	fmt.Println(StructName("id"))
	fmt.Println(StructName("abc_asdf"))
	fmt.Println(StructName("abc_asd123"))
	fmt.Println(StructName("foo_id"))
	fmt.Println(StructName("foo_Id"))
}

func TestSingular(t *testing.T) {
	fmt.Println(inflection.Singular("adsfasd_asdfsa"))
	fmt.Println(inflection.Singular("i am god"))
}
