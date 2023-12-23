//go:build wiondows

package executer

import "os/exec"

// TODO this function is not test in windows
func ExecuteCommand(args ...string) *exec.Cmd {
	wrapArgs := append([]string{"/C"}, args...)
	return exec.Command("cmd", wrapArgs...)
}
