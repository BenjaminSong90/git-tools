//go:build !wiondows

package executer

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(args ...string) *exec.Cmd {
	wrapArgs := append([]string{"-c"}, strings.Join(args, " "))
	return exec.Command("sh", wrapArgs...)
}
