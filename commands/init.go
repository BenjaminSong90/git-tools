package commands

import (
	"os"
	"path/filepath"

	"github.com/BenjaminSong90/git-tools/tools"
)

type InitCommand struct {
}

func (command *InitCommand) Execute(args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	gitExist, _, _ := FindDotGitFolder(pwd)

	if !gitExist {
		tools.Println("Please initialize git first! \n : git init", tools.Red)
		return nil
	}
	exist, realPath, deepth := FindDotGitToolsFolder(pwd)
	if exist && deepth == 0 {
		tools.Println("Reinitialized existing Git-tools repository in "+realPath, tools.Green)
		// TODO 添加再次初始化的逻辑
	} else {
		dotGitFolderName := filepath.Join(pwd, tools.GitToolsDirName)
		err := os.Mkdir(dotGitFolderName, 0755) // 0755 表示默认文件夹权限
		if err != nil {
			tools.Println("git tools init fail", tools.Red)
			return err
		}

		return InitGitToolsBranchInfoFile(dotGitFolderName)
	}

	return nil
}
