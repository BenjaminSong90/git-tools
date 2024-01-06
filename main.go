package main

import (
	"fmt"
	"os"

	"github.com/BenjaminSong90/git-tools/commands"
	"github.com/jessevdk/go-flags"
)

type Option struct {
	Init        commands.InitCommand        `command:"init"`
	Branch      commands.BranchCommand      `command:"branch"`
	Verify      commands.VerifyCommand      `command:"verify"`
	BranchGroup commands.BranchGroupCommand `command:"group"`
}

func main() {
	var opt Option
	opt.Branch.OnAttach()
	opt.BranchGroup.OnAttach()

	parser := flags.NewParser(&opt, flags.HelpFlag)
	var err error
	if len(os.Args) == 1 {
		_, err = parser.ParseArgs([]string{"--help"})
	} else {
		_, err = parser.Parse()
	}

	if err != nil {
		fmt.Println(err.Error())
	}

}
