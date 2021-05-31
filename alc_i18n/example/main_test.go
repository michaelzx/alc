package main

import (
	"github.com/michaelzx/alc/alc_i18n"
	"log"
	"path/filepath"
	"testing"
)

func init() {
	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	alc_i18n.Init(dir)
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alc_i18n.Msg(alc_i18n.ZhCN, "a.b").Trans()
		alc_i18n.Msg(alc_i18n.EnUS, "a.b").Trans()
		alc_i18n.Msg(alc_i18n.EnUS, "a.b.c").Trans()
		alc_i18n.Msg(alc_i18n.EnUS, "a").Trans()
		// 中文为key
		alc_i18n.Msg(alc_i18n.ZhCN, "提示.第一个").Trans()
		alc_i18n.Msg(alc_i18n.EnUS, "提示.第一个").Trans()
		// 解析变量
		alc_i18n.Msg(alc_i18n.ZhCN, "c").TransWithValues(map[string]interface{}{
			"Name": "Michael",
		})
		alc_i18n.Msg(alc_i18n.EnUS, "c").TransWithValues(map[string]interface{}{
			"Name": "Michael",
		})
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alc_i18n.Msg(alc_i18n.ZhCN, "a.b").Trans()
	}
}
func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alc_i18n.Msg(alc_i18n.EnUS, "required").Tag("validator").Trans()
	}
}
func Benchmark3(b *testing.B) {
	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	alc_i18n.Init(dir)
	for i := 0; i < b.N; i++ {
		alc_i18n.Msg(alc_i18n.EnUS, "c").TransWithValues(map[string]interface{}{
			"Name": "Michael",
		})
	}
}
