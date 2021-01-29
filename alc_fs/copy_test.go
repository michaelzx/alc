package alc_fs

import "testing"

func TestCopyDir(t *testing.T) {
	CopyDir("/Users/michael/workspace/cirs/cirs-doc-keeper/test-storage/disks/测试盘/test", "/Users/michael/workspace/cirs/cirs-doc-keeper/test-storage/disks/测试盘/新建文件夹")
}
