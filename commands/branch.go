package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BenjaminSong90/git-tools/executer"
	"github.com/BenjaminSong90/git-tools/tools"
)

type BranchCommand struct {
	Verify   func() `long:"verify"`
	Describe string `long:"describe" description:"branch describe" default:""`
	Name     string `long:"name" description:"branch name" default:""`

	hasCommandExec bool
}

func (command *BranchCommand) Execute(args []string) error {

	if command.Describe != "" || command.Name != "" {
		command.hasCommandExec = true
		setBranchDescribe(command.Name, command.Describe)
	}

	if !command.hasCommandExec {
		showBranchInfo()
	}

	return nil
}

func (command *BranchCommand) OnAttach() {
	command.Verify = func() {
		command.hasCommandExec = true
		branchVerify()
	}
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
		tools.Println("local branch info is error,please execute 'verify' \n error info : "+err.Error(), tools.Red)
		return
	}

	branches, err := getLocalBranches()
	if err != nil {
		tools.Println("read local branch error \n error info : "+err.Error(), tools.Red)
		return
	}

	bn := []string{}

	for _, b := range *branches {
		bn = append(bn, b.Name)
	}

	recoedBranchInfo.Verify(&bn)

	err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

	if err != nil {
		tools.Println("branch verify error : "+err.Error(), tools.Red)
	}
}

// 设置分支描述
func setBranchDescribe(name, desc string) {
	if name == "" && desc == "" {
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		tools.Println("cannot get path info, error : "+err.Error(), tools.Red)
		return
	}

	if !VerifyGitToolsEnv(wd) {
		return
	}

	_, folderPath, _ := FindDotGitToolsFolder(wd)

	recoedBranchInfo, err := LoadGitToolsBranchInfo(folderPath)

	if err != nil {
		tools.Println("local branch info is error,please execute 'verify' \n error info : "+err.Error(), tools.Red)
		return
	}

	setBranchName := name

	if setBranchName == "" {
		cbn, err := executer.GetCurrentBranch()
		if err != nil {
			return
		}
		setBranchName = cbn
	}

	recoedBranchInfo.SetBranchDesc(setBranchName, desc)

	err = WirteGitToolsBranchInfo(recoedBranchInfo, folderPath)

	if err != nil {
		tools.Println("branch set desc error : "+err.Error(), tools.Red)
	}
}

// 显示branch 信息
func showBranchInfo() {
	wd, err := os.Getwd()
	if err != nil {
		tools.Println("cannot get path info, error : "+err.Error(), tools.Red)
		return
	}

	exist, folderPath, _ := FindDotGitToolsFolder(wd)

	if !exist {
		return
	}

	recoedBranchInfo, err := LoadGitToolsBranchInfo(folderPath)

	if err != nil {
		tools.Println("local branch info is error,please execute 'verify' \n error info : "+err.Error(), tools.Red)
		return
	}

	showMap := make(map[string]interface{})
	showMap["branch"] = recoedBranchInfo.Branches
	showMap["group"] = recoedBranchInfo.BranchGroups

	jsonByte, err := json.MarshalIndent(showMap, "", " ")
	if err != nil {
		return
	}

	fmt.Println(string(jsonByte))
}
