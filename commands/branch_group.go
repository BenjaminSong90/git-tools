package commands

import "github.com/BenjaminSong90/git-tools/tools"

type BranchGroupCommand struct {
	CreateOwner           string   `long:"create" short:"c" description:"create group, need follow owner branch name"`
	Name                  string   `long:"name" short:"n" description:"group name"`
	Branches              []string `long:"branches" short:"b" description:"branch list for add or remove"`
	Description           string   `long:"description" short:"d" description:"group describe"`
	AddActionGroupName    string   `long:"add" short:"a" description:"add branch to group"`
	RemoveActionGroupName string   `long:"remove" description:"remove branch from group"`
}

func (command *BranchGroupCommand) Execute(args []string) error {
	command.create()
	command.setDesc()
	command.addBranch()
	command.deleteBranch()
	return nil
}

func (command *BranchGroupCommand) create() {
	if command.CreateOwner == "" {
		return
	}

	if command.Name == "" {
		tools.Println("group name cannot empty", tools.Red)
		return
	}

	branchInfo, gitToolsFolderPath, err := getValidBranchInfoAndPath()

	if err != nil {
		return
	}

	err = branchInfo.CreateGroup(command.Name, command.CreateOwner, command.Branches)

	if err != nil {
		tools.Println("create group fail, info: "+err.Error(), tools.Red)
		return
	}

	err = WirteGitToolsBranchInfo(branchInfo, gitToolsFolderPath)

	if err != nil {
		tools.Println("create group error : "+err.Error(), tools.Red)
	}

}

func (command *BranchGroupCommand) setDesc() {

	if command.Name == "" || command.Description == "" {
		return
	}

	branchInfo, gitToolsFolderPath, err := getValidBranchInfoAndPath()

	if err != nil {
		tools.Println("load branch info error : "+err.Error(), tools.Red)
		return
	}

	branchInfo.SetGroupDesc(command.Name, command.Description)

	err = WirteGitToolsBranchInfo(branchInfo, gitToolsFolderPath)

	if err != nil {
		tools.Println("set group desc error : "+err.Error(), tools.Red)
	}

}

func (command *BranchGroupCommand) addBranch() {
	if command.AddActionGroupName == "" {
		return
	}

	if len(command.Branches) == 0 {
		tools.Println("branch is empty", tools.Red)
		return
	}

	branchInfo, gitToolsFolderPath, err := getValidBranchInfoAndPath()

	if err != nil {
		tools.Println("load branch info error : "+err.Error(), tools.Red)
		return
	}

	branchInfo.AddBranchToGroup(command.AddActionGroupName, command.Branches)

	err = WirteGitToolsBranchInfo(branchInfo, gitToolsFolderPath)

	if err != nil {
		tools.Println("add branch error : "+err.Error(), tools.Red)
	}

}

func (command *BranchGroupCommand) deleteBranch() {
	if command.RemoveActionGroupName == "" {
		return
	}

	if len(command.Branches) == 0 {
		tools.Println("branch is empty", tools.Red)
		return
	}

	branchInfo, gitToolsFolderPath, err := getValidBranchInfoAndPath()

	if err != nil {
		tools.Println("load branch info error : "+err.Error(), tools.Red)
		return
	}

	branchInfo.RemoveBranchFromGroup(command.RemoveActionGroupName, command.Branches)

	err = WirteGitToolsBranchInfo(branchInfo, gitToolsFolderPath)

	if err != nil {
		tools.Println("add branch error : "+err.Error(), tools.Red)
	}

}
