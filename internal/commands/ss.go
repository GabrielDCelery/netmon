package commands

import (
	"context"
	"os/exec"
)

// Connection represents a single network connection from ss output.
type Connection struct {
	Protocol string
	State    string
	RecvQ    string
	SendQ    string
	Local    string
	Peer     string
	Process  string
}

// SSCommand executes the ss command to retrieve socket statistics.
type SSCommand struct {
}

// NewSSCommand creates a new SSCommand.
func NewSSCommand() *SSCommand {
	return &SSCommand{}
}

// Run executes "ss -tunap" and returns the raw output.
func (r *SSCommand) Run(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "ss", "-tunap")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Command returns the full command string being executed.
func (r *SSCommand) PrintCommandAsStr() string {
	return "ss -tunap"
}
