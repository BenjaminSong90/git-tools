package commands

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/BenjaminSong90/git-tools/data"
	"github.com/BenjaminSong90/git-tools/executer"
	"github.com/BenjaminSong90/git-tools/tools"
)

type InitCommand struct {
}

func (command *InitCommand) Execute(args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	gitExist, _ := FindDotGitFolder(pwd)

	if !gitExist {
		tools.Println("Please initialize git first! \n : git init", tools.Red)
		return nil
	}
	exist, realPath := FindDotGitToolsFolder(pwd)
	if exist {
		tools.Println("Reinitialized existing Git-tools repository in "+realPath, tools.Green)
		// TODO 添加再次初始化的逻辑
	} else {
		dotGitFolderName := filepath.Join(pwd, tools.GitToolsDirName)
		err := os.Mkdir(dotGitFolderName, 0755) // 0755 表示默认文件夹权限
		if err != nil {
			tools.Println("git tools init fail", tools.Red)
			return err
		}

		return initGitToolsFiles(dotGitFolderName)
	}

	return nil
}

// 初始化 .git-tools 文件夹下的文件
func initGitToolsFiles(rootDirPath string) error {
	newFile, err := tools.CreateFileByPath(filepath.Join(rootDirPath, tools.BranchInfoFileName))
	if err != nil {
		return nil
	}
	defer newFile.Close()

	localBranches, err := executer.GetLocalAllBranch()

	if err != nil {
		return err
	}

	branches := []data.Branch{}
	for _, b := range localBranches {
		bd := data.Branch{
			Name: b,
		}
		branches = append(branches, bd)
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
