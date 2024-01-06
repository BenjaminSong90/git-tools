package commands

import (
	"os"

	"github.com/BenjaminSong90/git-tools/executer"
	"github.com/BenjaminSong90/git-tools/tools"
)

type BranchCommand struct {
	Verify   func() `long:"verify"`
	Describe string `long:"describe"`
	Name     string `long:"name" description:"branch name" default:""`
}

func (command *BranchCommand) Execute(args []string) error {
	command.setBranchDescribe(command.Name, command.Describe)
	return nil
}

func (command *BranchCommand) OnAttach() {
	command.Verify = branchVerify
}

// 分支检查
func branchVerify() {
	wd, err := os.Getwd()
	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
		return
	}

	if !VerifyGitToolsEnv(wd) {
		return
	}

	_, folderPath, _ := FindDotGitToolsFolder(wd)

	recoedBranchInfo, err := LoadGitToolsBranchInfo(folderPath)

	if err != nil {
		tools.Println("git-tools info is error,please execute 'verify' \n error info : "+err.Error(), tools.Red)
		return
	}

	branches, err := getLocalBranches()
	if err != nil {
		tools.Println("read local branch error \n error info : "+err.Error(), tools.Red)
		return
	}

	bn := []string{}

	for _, b := range *branches {
		bn = append(bn, b.Name)
	}

	recoedBranchInfo.Verify(&bn)

	err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
	}
}

// 设置分支描述
func (command *BranchCommand) setBranchDescribe(name, desc string) {
	if name == "" && desc == "" {
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		tools.Println("cannot get current branch info, error : "+err.Error(), tools.Red)
		return
	}

	if !VerifyGitToolsEnv(wd) {
		return
	}

	_, folderPath, _ := FindDotGitToolsFolder(wd)

	recoedBranchInfo, err := LoadGitToolsBranchInfo(folderPath)

	if err != nil {
		tools.Println("git-tools info is error,please execute 'verify' \n error info : "+err.Error(), tools.Red)
		return
	}

	setBranchName := name

	if setBranchName == "" {
		cbn, err := executer.GetCurrentBranch()
		if err != nil {
			return
		}
		setBranchName = cbn
	}

	recoedBranchInfo.SetBranchDesc(setBranchName, desc)

	err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

	if err != nil {
		tools.Println("branch set desc error : "+err.Error(), tools.Red)
	}
}
