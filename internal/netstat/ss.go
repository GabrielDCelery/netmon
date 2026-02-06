package netstat

import (
	"context"
	"os/exec"
)

// Runner executes a system command and returns its output.
type Runner interface {
	Run(ctx context.Context) (string, error)
	Command() string
}

// SSRunner executes the ss command to retrieve socket statistics.
type SSRunner struct{}

// NewSSRunner creates a new SSRunner.
func NewSSRunner() *SSRunner {
	return &SSRunner{}
}

// Run executes "ss -tunap" and returns the raw output.
func (r *SSRunner) Run(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "ss", "-tunap")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Command returns the full command string being executed.
func (r *SSRunner) Command() string {
	return "ss -tunap"
}
