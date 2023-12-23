package commands

type InitCommand struct {
}

func (command *InitCommand) Execute(args []string) error {
	// fmt.Print("init command execute..... \n" + strings.Join(args, " "))
	// resultByte, err := executer.ExecuteCommand("git branch").Output()
	// if err != nil {
	// 	return err
	// }
	// fmt.Print(">>>>>>>: \n" + string(resultByte) + "<<<<<<<< \n")
	return nil
}
