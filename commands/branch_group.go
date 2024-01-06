package commands

import "github.com/BenjaminSong90/git-tools/tools"

type BranchGroupCommand struct {
	Owner       string   `long:"create" short:"c" description:"create group, need follow owner branch name"`
	Name        string   `long:"name" short:"n" description:"group name"`
	Branches    []string `long:"branches" short:"b" description:"branch list for add or remove"`
	Description string   `long:"description" short:"d" description:"group describe"`
}

func (command *BranchGroupCommand) Execute(args []string) error {
	command.create(command.Name, command.Owner, command.Branches)
	command.setDesc(command.Name, command.Description)
	return nil
}

func (command *BranchGroupCommand) OnAttach() {

}

func (command *BranchGroupCommand) create(name string, ownerName string, branches []string) {
	if ownerName == "" {
		return
	}

	if name == "" {
		tools.Println("group name cannot empty", tools.Red)
		return
	}

	branchInfo, gitToolsFolderPath, err := getValidBranchInfoAndPath()

	if err != nil {
		return
	}

	err = branchInfo.CreateGroup(name, ownerName, branches)

	if err != nil {
		tools.Println("create group fail, info: "+err.Error(), tools.Red)
		return
	}

	err = WirteGitToolsBranchInfo(branchInfo, gitToolsFolderPath)

	if err != nil {
		tools.Println("create group error : "+err.Error(), tools.Red)
	}

}

func (command *BranchGroupCommand) setDesc(name string, desc string) {

	if name == "" || desc == "" {
		return
	}

	branchInfo, gitToolsFolderPath, err := getValidBranchInfoAndPath()

	if err != nil {
		return
	}

	branchInfo.SetGroupDesc(name, desc)

	err = WirteGitToolsBranchInfo(branchInfo, gitToolsFolderPath)

	if err != nil {
		tools.Println("set group desc error : "+err.Error(), tools.Red)
	}

}
