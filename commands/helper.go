package commands

// 通用的helper

import "github.com/BenjaminSong90/git-tools/tools"

//查找 .git-tools 文件夹和子文件工具类

func FindDotGitToolsFolder(path string) (exist bool, realPath string) {
	exist, realPath = tools.IsFileExistsAlongPath(path, tools.GitToolsDirName)

	return
}

func FindDotGitFolder(path string) (exist bool, realPath string) {
	exist, realPath = tools.IsFileExistsAlongPath(path, ".git")

	return
}
