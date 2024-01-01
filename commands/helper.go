package commands

// 通用的helper

import (
	"encoding/json"
	"path/filepath"

	"github.com/BenjaminSong90/git-tools/data"
	"github.com/BenjaminSong90/git-tools/executer"
	"github.com/BenjaminSong90/git-tools/tools"
)

//查找 .git-tools 文件夹和子文件工具类

func FindDotGitToolsFolder(path string) (exist bool, realPath string, deepth int) {
	exist, realPath, deepth = tools.IsFileExistsAlongPath(path, tools.GitToolsDirName)

	return
}

// 查找 .git 文件夹 用于检查git是否初始化
func FindDotGitFolder(path string) (exist bool, realPath string, deepth int) {
	exist, realPath, deepth = tools.IsFileExistsAlongPath(path, ".git")

	return
}

// 加载git-tools branchInfo 数据
func LoadGitToolsBranchInfo(rootPath string) (*data.BranchInfo, error) {
	branchInfo, err := tools.LoadFileJsonData[data.BranchInfo](filepath.Join(rootPath, tools.BranchInfoFileName))

	return branchInfo, err

}

// 初始化 .git-tools 文件夹下的文件
func InitGitToolsBranchInfoFile(rootDirPath string) error {
	newFile, err := tools.CreateFileByPath(filepath.Join(rootDirPath, tools.BranchInfoFileName))
	if err != nil {
		return nil
	}
	defer newFile.Close()

	branches, err := getLocalBranches()
	if err != nil {
		return err
	}

	branchInfo := data.BranchInfo{
		Version:      tools.BranchInfoVersion,
		Branches:     branches,
		BranchGroups: []data.BranchGroup{},
	}
	// 将结构体转换为 JSON 字符串
	jsonByte, err := json.Marshal(branchInfo)
	if err != nil {
		return nil
	}

	_, err = newFile.Write(jsonByte)
	return err
}

// 获取本地的git branch 信息
func getLocalBranches() ([]data.Branch, error) {
	localBranches, err := executer.GetLocalAllBranch()

	if err != nil {
		return nil, err
	}
	if len(localBranches) == 0 {
		emtpy := []data.Branch{}
		return emtpy, nil
	}
	branches := []data.Branch{}
	for _, b := range localBranches {
		bd := data.Branch{
			Name: b,
		}
		branches = append(branches, bd)
	}

	return branches, nil
}
