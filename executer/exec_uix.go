//go:build !wiondows

package executer

import "os/exec"

func ExecuteCommand(args ...string) *exec.Cmd {
	wrapArgs := append([]string{"-c"}, args...)
	return exec.Command("sh", wrapArgs...)
}
