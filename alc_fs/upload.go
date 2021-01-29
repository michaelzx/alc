package alc_fs

import (
	"alchemy/alc/alc_errs"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"time"
)

type UploadResult struct {
	OriginalFilename  string
	OriginalExtension string
	ServerPath        string // 相对于应用的服务器路径
}

// Upload 上传文件
// @savePath 相对于应用的目录
// @fullPath 最后保存的完整路径
func Upload(gc *gin.Context, postFormName string, saveDir string, allowTypes ...FileType) (result *UploadResult, err error) {
	var form *multipart.Form
	form, err = gc.MultipartForm()
	if err != nil {
		return
	}
	files := form.File[postFormName]
	if len(files) == 0 {
		err = errors.New("请选择要上传的文件")
		return
	}
	if len(files) != 1 {
		err = errors.New("仅支持单文件上传")
		return
	}
	monthDir := time.Now().Format("200601")
	savePath := filepath.Join(AppPath, saveDir, monthDir)
	// 如果目录不存在，则创建
	if !Exists(savePath) {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			err = alc_errs.Wrap(err, "目录创建错误")
			return
		}
	}
	// 检查文件类型
	err = checkFileType(files, allowTypes...)
	if err != nil {
		return
	}
	file := files[0]
	fileName := RandomFilename() + path.Ext(file.Filename)
	serverPath := filepath.Join(savePath, fileName)

	err = gc.SaveUploadedFile(file, serverPath)
	if err != nil {
		return
	}
	result = &UploadResult{
		OriginalFilename:  file.Filename,
		OriginalExtension: path.Ext(file.Filename),
		ServerPath:        serverPath,
	}
	return
}
