package base

import "os"

// FileIsExist 判断文件是否存在
func FileIsExist(filepath string) bool {
	f, err := os.Open(filepath)
	defer f.Close()
	return os.IsNotExist(err)
}
