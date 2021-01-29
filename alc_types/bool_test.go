package alc_types

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BoolTest struct {
	B Bool
}

func TestBoolFromJson(t *testing.T) {
	str := `{"B":false}`
	boolTest := BoolTest{}
	err := json.Unmarshal([]byte(str), &boolTest)
	if err != nil {
		println(err)
	}
	fmt.Printf("%#v\n", boolTest)
	if boolTest.B {
		fmt.Println("B is true")
	} else {
		fmt.Println("B is false")
	}
}

func TestBoolToJson(t *testing.T) {
	boolTest := BoolTest{
		B: true,
	}
	jsonStr, err := json.Marshal(boolTest)
	if err != nil {
		println(err)
	}
	fmt.Println(string(jsonStr))
}
