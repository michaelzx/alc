package alc_fs

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const WebDir = "/web"
const ResourceDir = "/resource"
const UploadsUrl = "/uploads"
const ThemeUrl = "/theme"

var AppPath = PathAppRunning()
var WebPath = filepath.Join(AppPath, WebDir)
var ResourcePath = filepath.Join(AppPath, ResourceDir)
var UploadsPath = filepath.Join(WebPath, UploadsUrl)
var ThemePath = filepath.Join(WebPath, ThemeUrl)

// PathAppRunning 当前程序运行物理路径
func PathAppRunning() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// CreateIfNotExist 目录如果不存在则创建
func CreateIfNotExist(path string) {
	if !Exists(path) {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatal(err)
		}
	}
}
