package commands

import (
	"fmt"
	"strings"
)

type BranchCommand struct {
}

func (command *BranchCommand) Execute(args []string) error {
	fmt.Print("branch command execute....." + strings.Join(args, " "))
	return nil
}
