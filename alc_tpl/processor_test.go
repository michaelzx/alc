package pld_tpl

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	println(regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`).MatchString(`<img :src="asdfsa">`))           // false
	println(regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`).MatchString(`<img src="asdfsa">`))            // true
	println(regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`).MatchString(`<img  src="asdfsa">`))           // true
	println(regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`).MatchString(`<img  :src="asdfsa">`))          // false
	println(regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`).MatchString(`<img alt="123" src="asdfsa">`))  // true
	println(regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`).MatchString(`<img alt="123" :src="asdfsa">`)) // false
}
