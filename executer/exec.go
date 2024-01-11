package executer

import (
	"strings"
)

// 获取本地的分支
func GetLocalAllBranch() ([]string, error) {
	resultByte, err := ExecuteCommand("git", "branch").Output()
	if err != nil {
		return nil, err
	}
	resultStr := strings.TrimSpace(string(resultByte))

	localBranhes := []string{}

	if len(resultStr) == 0 {
		return localBranhes, nil
	}
	currentBranches := strings.Split(resultStr, "\n")

	for _, branch := range currentBranches {
		if strings.HasPrefix(branch, "*") {
			cb := strings.TrimSpace(strings.Replace(branch, "*", "", 1))
			localBranhes = append(localBranhes, cb)
		} else {
			cb := strings.TrimSpace(branch)
			localBranhes = append(localBranhes, cb)
		}
	}

	return localBranhes, nil
}

// 获取当前分支
func GetCurrentBranch() (string, error) {
	resultByte, err := ExecuteCommand("git", "branch", "--show-current").Output()
	if err != nil {
		return "", err
	}
	resultStr := strings.TrimSpace(string(resultByte))

	return resultStr, nil
}
