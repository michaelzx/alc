package alc_fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_randomFilename(t *testing.T) {
	fmt.Println(RandomFilename())
}
func TestGetFileType(t *testing.T) {
	getHex(t, "/Users/michael/Downloads/xvideos.com_6e202bc9af8d7efd632aa1a9e82081ac.mp4")
}
func getHex(t *testing.T, fp string) {
	// f, err := os.Open("C:\\Users\\Administrator\\Desktop\\api.html")
	f, err := os.Open(fp)
	if err != nil {
		t.Logf("open error: %v", err)
	}

	fSrc, err := ioutil.ReadAll(f)
	fileCode := bytesToHexString(fSrc[:20])
	t.Log(fileCode)
}
