package alc_fs

import (
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	// f1, err := os.Open("/home/zzw/test_data/ziptest/gophercolor16x16.png")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer f1.Close()
	// f2, err := os.Open("/home/zzw/test_data/ziptest/readme.notzip")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer f2.Close()
	f3, err := os.Open("/Users/michael/Downloads/alibaba-dingtalk-service-sdk-1.0.1-sources (2)")
	if err != nil {
		t.Fatal(err)
	}
	defer f3.Close()
	var files = []*os.File{f3}
	dest := "/Users/michael/Downloads/test.zip"
	err = Compress(files, dest)
	if err != nil {
		t.Fatal(err)
	}
}
func TestDeCompress(t *testing.T) {
	err := DeCompress("/home/zzw/test_data/test.zip", "/home/zzw/test_data/de")
	if err != nil {
		t.Fatal(err)
	}
}
