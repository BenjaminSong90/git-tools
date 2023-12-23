package commands

import (
	"fmt"
	"strings"
)

type VerifyCommand struct {
}

func (command *VerifyCommand) Execute(args []string) error {
	fmt.Print("verify command execute....." + strings.Join(args, " "))
	return nil
}
