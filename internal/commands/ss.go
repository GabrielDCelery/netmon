package commands

import (
	"context"
	"os/exec"
	"strings"
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
type SSCommand struct{}

// NewSSCommand creates a new SSCommand.
func NewSSCommand() *SSCommand {
	return &SSCommand{}
}

// Run executes "ss -tunap" and returns the raw output.
func (r *SSCommand) Run(ctx context.Context) ([]Connection, error) {
	cmd := exec.CommandContext(ctx, "ss", "-tunap")
	out, err := cmd.Output()
	if err != nil {
		return []Connection{}, err
	}
	return Parse(string(out)), nil
}

// Command returns the full command string being executed.
func (r *SSCommand) PrintCommandAsStr() string {
	return "ss -tunap"
}

// Parse takes raw ss -tunap output and returns a slice of Connection structs.
func Parse(raw string) []Connection {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	if len(lines) < 2 {
		return nil
	}

	var connections []Connection

	// Skip header line
	for _, line := range lines[1:] {
		conn, ok := parseLine(line)
		if ok {
			connections = append(connections, conn)
		}
	}

	return connections
}

func parseLine(line string) (Connection, bool) {
	fields := strings.Fields(line)
	if len(fields) < 6 {
		return Connection{}, false
	}

	conn := Connection{
		Protocol: fields[0],
		State:    fields[1],
		RecvQ:    fields[2],
		SendQ:    fields[3],
		Local:    fields[4],
		Peer:     fields[5],
	}

	if len(fields) > 6 {
		conn.Process = strings.Join(fields[6:], " ")
	}

	return conn, true
}
