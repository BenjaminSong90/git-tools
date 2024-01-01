package commands

import (
	"os"
	"path/filepath"

	"github.com/BenjaminSong90/git-tools/tools"
)

type VerifyCommand struct {
}

func (command *VerifyCommand) Execute(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	exist, folderPath, _ := FindDotGitToolsFolder(wd)
	if err != nil {
		return err
	}

	if !exist {
		tools.Println("Please initialize git-tools first! \n git-tools init", tools.Red)
		return nil
	}

	_, err = LoadGitToolsBranchInfo(folderPath)
	if err != nil {
		branchInfoPath := filepath.Join(folderPath, tools.BranchInfoFileName)

		err := os.Remove(branchInfoPath)
		if err != nil {
			return err
		}
		return InitGitToolsBranchInfoFile(folderPath)
	}

	return nil
}
