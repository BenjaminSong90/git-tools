package commands

import (
	"os"

	"github.com/BenjaminSong90/git-tools/tools"
)

type BranchCommand struct {
	Verify func() `long:"verify"`
	Op     string `long:"op" description:"operation to execute" default:"default"`
	Op2    string `long:"op2" description:"operation to execute"`
}

func (command *BranchCommand) Execute(args []string) error {
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

	recoedBranchInfo.Verify(branches)

	err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
	}
}
