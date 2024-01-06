package commands

// 通用的helper

import (
	"errors"
	"os"
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

func WirteGitToolsBranchInfo(branchInfo *data.BranchInfo, rootPath string) error {
	return tools.WriteJsonData2File[data.BranchInfo](branchInfo, filepath.Join(rootPath, tools.BranchInfoFileName))
}

// 初始化 .git-tools 文件夹下的文件
func InitGitToolsBranchInfoFile(rootDirPath string) error {

	branches, err := getLocalBranches()
	if err != nil {
		return err
	}

	branchInfo := data.BranchInfo{
		Version:      tools.BranchInfoVersion,
		Branches:     *branches,
		BranchGroups: []data.BranchGroup{},
	}

	return tools.WriteJsonData2File(&branchInfo, filepath.Join(rootDirPath, tools.BranchInfoFileName))
}

// 获取本地的git branch 信息
func getLocalBranches() (*[]data.Branch, error) {
	localBranches, err := executer.GetLocalAllBranch()

	if err != nil {
		return nil, err
	}

	branches := []data.Branch{}
	if len(localBranches) == 0 {
		return &branches, nil
	}

	for _, b := range localBranches {
		bd := data.Branch{
			Name: b,
		}
		branches = append(branches, bd)
	}

	return &branches, nil
}

// 检查git 和 git是否初始化
func VerifyGitToolsEnv(path string) bool {
	gitExist, _, _ := FindDotGitFolder(path)
	if !gitExist {
		tools.Println("Please initialize git first! \n : git init", tools.Red)
		return false
	}
	gitToolsExits, _, _ := FindDotGitToolsFolder(path)
	if !gitToolsExits {
		tools.Println("Please initialize git-tools first! \n git-tools init", tools.Red)
	}
	return gitToolsExits
}

// 获取 branchinfo 有效的数据 和 路径
func getValidBranchInfoAndPath() (branchInfo *data.BranchInfo, path string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	if !VerifyGitToolsEnv(wd) {
		return nil, "", errors.New("env is error")
	}

	_, path, _ = FindDotGitToolsFolder(wd)

	branchInfo, err = LoadGitToolsBranchInfo(path)

	return
}
