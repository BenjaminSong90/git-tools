package tools

import (
	"os"
	"path/filepath"
)

// 检查是否是合法的路径
func IsPathValid(path string) bool {
	// 检查路径是否为绝对路径
	isAbsolute := filepath.IsAbs(path)

	// 清理和规范化路径
	cleanPath := filepath.Clean(path)

	// 检查清理后的路径和原始路径是否相同
	isValid := cleanPath == path && isAbsolute

	return isValid
}

// 路径是否存在
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // 路径存在
	}
	if os.IsNotExist(err) {
		return false // 路径不存在
	}
	return false // 其他错误，假定路径不存在
}

// 检查文件是否存在路径目录中
func IsFileExistsAlongPath(dirPath string, fileName string) (bool, string) {
	dir := dirPath

	for {
		fullPath := filepath.Join(dir, fileName)

		// 检查文件是否存在
		_, err := os.Stat(fullPath)
		if err == nil {
			return true, fullPath // 文件存在
		}
		if !os.IsNotExist(err) {
			return false, "" // 其他错误，假定文件不存在
		}

		// 获取上级目录
		parent := filepath.Dir(dir)
		if parent == dir {
			break // 到达根目录，退出循环
		}
		dir = parent
	}

	return false, "" // 文件不存在
}

// 通过文件路径创建文件
func CreateFileByPath(filePath string) (*os.File, error) {
	newFile, err := os.Create(filePath)

	return newFile, err
}
