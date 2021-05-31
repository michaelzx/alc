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
	fmt.Println(alc_i18n.Msg(alc_i18n.ZhCN, "a.b").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EnUS, "a.b").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EnUS, "a.b.c").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EnUS, "a").Trans())
	// 中文为key
	fmt.Println(alc_i18n.Msg(alc_i18n.ZhCN, "提示.第一个").Trans())
	fmt.Println(alc_i18n.Msg(alc_i18n.EnUS, "提示.第一个").Trans())
	// 解析变量
	fmt.Println(alc_i18n.Msg(alc_i18n.ZhCN, "c").TransWithValues(map[string]interface{}{
		"Name": "Michael",
	}))
	fmt.Println(alc_i18n.Msg(alc_i18n.EnUS, "c").TransWithValues(map[string]interface{}{
		"Name": "Michael",
	}))
}
