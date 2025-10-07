package git

import (
	"os/exec"
)

func GitCommands(fs *flag.FlagSet, args []string) error {
	switch 
}
	return 

func RunGitCommand(args ...string) (string, error) {
	cmd := exec.Command ("git", args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}