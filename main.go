package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/**
merge checker 根据配置来检查配置的分支是否已经merge到当前的分支

对应的配置信息结果：
{
	"branch":[
		"branch_a",
		"branch_b"
	]
}

*/

type Config struct {
	Branch []string `json:"branch"`
}

func main() {
	var configFilePath string //merge检查confige文案路径

	flag.StringVar(&configFilePath, "config_path", "./merge_checker.json", "merge检查confige文案路径")

	flag.Parse()

	fmt.Printf("configFilePath: %s \n\n", configFilePath)

	fi, err := os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s 文件不存在 \n", configFilePath)
		} else {
			fmt.Printf("%s 文件打开错误 err:%s \n", configFilePath, err)
		}
		return
	}

	if fi.IsDir() {
		fmt.Printf("%s 不是文件 \n", configFilePath)
		return
	}

	configContent, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("%s 文件打开失败 \n", err)
		return
	}

	var config Config
	err = json.Unmarshal(configContent, &config)
	if err != nil {
		fmt.Printf("%s json 序列化失败 err: %s \n", configFilePath, err)
		return
	}

	resultByte, err := exec.Command("sh", "-c", "git branch --merged").Output()

	if err != nil {
		fmt.Printf("错误: %s \n", err)
		return
	}
	resultStr := strings.TrimSpace(string(resultByte))
	currentBranchArray := strings.Split(resultStr, "\n")

	localBranchSet := make(map[string]bool)
	for _, branch := range currentBranchArray {
		if strings.HasPrefix(branch, "*") {
			cb := strings.TrimSpace(strings.Replace(branch, "*", "", 1))
			localBranchSet[strings.ToLower(cb)] = true
		} else {
			cb := strings.TrimSpace(branch)
			localBranchSet[strings.ToLower(cb)] = true
		}
	}

	var notMergeArr []string
	for _, branch := range config.Branch {
		if _, ok := localBranchSet[strings.ToLower(branch)]; !ok {
			notMergeArr = append(notMergeArr, branch)
		}
	}

	if len(notMergeArr) != 0 {
		fmt.Printf("\033[1;31m 没有merge的分支 : %s\033[0m \n\n", strings.Join(notMergeArr, ", "))
	} else {
		fmt.Print("\033[1;32m 所有分支已经merge \033[0m \n\n\n")
	}

}
