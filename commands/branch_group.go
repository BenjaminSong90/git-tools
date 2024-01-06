package commands

import "github.com/BenjaminSong90/git-tools/tools"

type BranchGroupCommand struct {
	Owner    string   `long:"create" description:"create group, need follow owner branch name"`
	Name     string   `long:"name" description:"group name"`
	Branches []string `long:"branches" description:"branch list for add or remove"`
}

func (command *BranchGroupCommand) Execute(args []string) error {
	command.create(command.Name, command.Owner, command.Branches)
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
