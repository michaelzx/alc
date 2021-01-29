package alc_fs

import (
	"alchemy/alc/alc_errs"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

type SaveFormFileResult struct {
	OriginalFilename  string
	OriginalExtension string
	ServerPath        string // 相对于应用的服务器路径
}

// SaveFormFile 保存form文件
func SaveFormFile(formFile *multipart.FileHeader, saveDir string, fileName string) (result *SaveFormFileResult, err error) {
	if formFile == nil {
		err = errors.New("请选择要上传的文件")
		return
	}
	saveDirPath := filepath.Join(AppPath, saveDir)
	// 如果目录不存在，则创建
	if !Exists(saveDirPath) {
		err = os.MkdirAll(saveDirPath, os.ModePerm)
		if err != nil {
			err = alc_errs.Wrap(err, "目录创建错误")
			return
		}
	}
	serverFilePath := filepath.Join(saveDirPath, fileName)

	src, err := formFile.Open()
	if err != nil {
		return
	}
	defer src.Close()

	out, err := os.Create(serverFilePath)
	if err != nil {
		return
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return
	}

	result = &SaveFormFileResult{
		OriginalFilename:  formFile.Filename,
		OriginalExtension: path.Ext(formFile.Filename),
		ServerPath:        serverFilePath,
	}
	return
}
