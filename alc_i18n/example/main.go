package main

import (
	"fmt"
	"github.com/michaelzx/alc/alc_i18n"
	"log"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	alc_i18n.Init(dir)
	fmt.Println(alc_i18n.Msg(alc_i18n.ZhCn, "a.b").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EN, "a.b").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EN, "a.b.c").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EN, "a").Trans())
	// 中文为key
	fmt.Println(alc_i18n.Msg(alc_i18n.ZhCn, "提示.第一个").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EN, "提示.第一个").Trans())
	// 解析变量
	fmt.Println(alc_i18n.Msg(alc_i18n.ZhCn, "c").TransWithValues(map[string]interface{}{
		"Name": "Michael",
	}))
	fmt.Println(alc_i18n.Msg(alc_i18n.EN, "c").TransWithValues(map[string]interface{}{
		"Name": "Michael",
	}))
}
