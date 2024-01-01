package commands

import (
	"os"

	"github.com/BenjaminSong90/git-tools/data"
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

func branchVerify() {
	wd, err := os.Getwd()
	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
		return
	}

	exist, folderPath, _ := FindDotGitToolsFolder(wd)
	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
		return
	}

	if !exist {
		tools.Println("Please initialize git-tools first! \n git-tools init", tools.Red)
		return
	}

	recoedBranchInfo, err := LoadGitToolsBranchInfo(folderPath)

	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
		return
	}

	branches, err := getLocalBranches()
	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
		return
	}

	if len(branches) == 0 {
		recoedBranchInfo.BranchGroups = []data.BranchGroup{}
		recoedBranchInfo.Branches = []data.Branch{}
		err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

		if err != nil {
			tools.Println("branch verify error : "+err.Error(), tools.Red)
		}
	} else {
		removeAndAddBranches(recoedBranchInfo, branches)

		err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

		if err != nil {
			tools.Println("branch verify error : "+err.Error(), tools.Red)
		}

		//TODO 处理branch分组的数据
	}
}

// 与本地的分支进行比较，移除已经删除的分支，添加新的分支
func removeAndAddBranches(branchInfo *data.BranchInfo, localBranches []data.Branch) {

	recordBranchMap := make(map[string]data.Branch)

	if len(localBranches) == 0 {
		branchInfo.Branches = []data.Branch{}
		return
	}

	for _, b := range branchInfo.Branches {
		recordBranchMap[b.Name] = b
	}

	notDeletedBranches := []data.Branch{}

	for _, b := range localBranches {
		if ndb, ok := recordBranchMap[b.Name]; ok {
			notDeletedBranches = append(notDeletedBranches, ndb)
		} else {
			notDeletedBranches = append(notDeletedBranches, b)
		}
	}

	branchInfo.Branches = notDeletedBranches

}
