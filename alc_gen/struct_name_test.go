package alc_gen

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"testing"
)

func TestStructName(t *testing.T) {
	fmt.Println(structName("_"))
	fmt.Println(structName("-"))
	fmt.Println(structName("id"))
	fmt.Println(structName("abc_asdf"))
	fmt.Println(structName("abc_asd123"))
	fmt.Println(structName("foo_id"))
	fmt.Println(structName("foo_Id"))
}

func TestSingular(t *testing.T) {
	fmt.Println(inflection.Singular("adsfasd_asdfsa"))
	fmt.Println(inflection.Singular("i am god"))
}
