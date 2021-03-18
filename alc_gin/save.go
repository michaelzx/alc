package alc_gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/alc/alc_errs"
	"github.com/michaelzx/alc/alc_fs"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

type SaveResult struct {
	Filename   string
	Extension  string
	ServerPath string `json:"-"` // 相对于应用的服务器路径
	WebPath    string // 相对于web的路径
}

// Save 保存文件
// @savePath 相对于应用的目录
// @fullPath 最后保存的完整路径
func Save(gc *gin.Context, postFormName string, webDir, saveName string, allowTypes ...alc_fs.FileType) (result *SaveResult, err error) {
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
	savePath := filepath.Join(alc_fs.WebPath, webDir)
	// 如果目录不存在，则创建
	if !alc_fs.Exists(savePath) {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			err = alc_errs.Wrap(err, "目录创建错误")
			return
		}
	}
	// 检查文件类型
	if len(allowTypes) > 0 {
		err = alc_fs.CheckFileType(files, allowTypes...)
		if err != nil {
			return
		}
	}
	file := files[0]
	fileName := saveName + path.Ext(file.Filename)
	serverPath := filepath.Join(savePath, fileName)
	webPath := filepath.Join(webDir, fileName)

	err = gc.SaveUploadedFile(file, serverPath)
	if err != nil {
		return
	}
	result = &SaveResult{
		Filename:   fileName,
		Extension:  path.Ext(fileName),
		ServerPath: serverPath,
		WebPath:    webPath,
	}
	return
}
