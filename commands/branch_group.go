package commands

import (
	"fmt"
	"strings"
)

type BranchGroupCommand struct {
	Add    string `long:"add" description:"operation to execute"`
	Args   []string
	Result int64
}

func (command *BranchGroupCommand) Execute(args []string) error {
	fmt.Print("branch group command execute....." + strings.Join(args, " "))
	return nil
}
