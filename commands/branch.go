package commands

import (
	"os"

	"github.com/BenjaminSong90/git-tools/executer"
	"github.com/BenjaminSong90/git-tools/tools"
)

type BranchCommand struct {
	Verify   func()       `long:"verify"`
	Describe func(string) `long:"describe"`
	Op       string       `long:"op" description:"operation to execute" default:"default"`
	Op2      string       `long:"op2" description:"operation to execute"`
}

func (command *BranchCommand) Execute(args []string) error {
	return nil
}

func (command *BranchCommand) OnAttach() {
	command.Verify = branchVerify
	command.Describe = describe

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

func describe(desc string) {
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

	cbn, err := executer.GetCurrentBranch()
	if err != nil {
		return
	}
	recoedBranchInfo.SetBranchDesc(cbn, desc)

	err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
	}
}
