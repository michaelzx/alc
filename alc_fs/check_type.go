package alc_fs

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
)

func CheckFileType(files []*multipart.FileHeader, allowTypes ...FileType) error {
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return err
		}
		fSrc, err := ioutil.ReadAll(f)
		_ = f.Close()
		fSrc20 := fSrc[:20]
		ft := GetFileType(fSrc20)
		if ft == "" {
			switch {
			case FileIsMp4(fSrc20):
				ft = string(FileTypeMp4)
			case FileIsMp3(fSrc20):
				ft = string(FileTypeMp3)
			case FileIsWav(fSrc20):
				ft = string(FileTypeWav)
			}
		}
		allow := false
		for _, at := range allowTypes {
			if ft == string(at) {
				allow = true
				break
			}
		}
		if !allow {
			if ft == "" {
				return errors.New("无法识别该文件类型，请联系开发人员")
			} else {
				return errors.New("不允许上传 " + ft + " 类型的文件")
			}
		}
	}
	return nil
}
