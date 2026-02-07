package tui

import (
	"context"
	"strings"

	"github.com/GabrielDCelery/netmon/internal/commands"
	"github.com/charmbracelet/bubbles/table"
)

type CommandRunner interface {
	Run(ctx context.Context) CommandRunResults
	Columns() []table.Column
	Rows() []table.Row
	PrintCommandAsStr() string
}

type CommandRunResults error

// SSCommandRunner runs the SS Command
type SSCommandRunner struct {
	command     commands.Command
	connections []commands.Connection
}

// NewSSCommandRunner creates a new SSCommandRunner
func NewSSCommandRunner(command *commands.SSCommand) *SSCommandRunner {
	return &SSCommandRunner{
		command:     command,
		connections: []commands.Connection{},
	}
}

func (r *SSCommandRunner) Columns() []table.Column {
	return []table.Column{
		{Title: "Proto", Width: 6},
		{Title: "State", Width: 12},
		{Title: "Recv-Q", Width: 8},
		{Title: "Send-Q", Width: 8},
		{Title: "Local Address", Width: 25},
		{Title: "Peer Address", Width: 25},
		{Title: "Process", Width: 30},
	}
}

func (r *SSCommandRunner) Rows() []table.Row {
	rows := make([]table.Row, len(r.connections))
	for i, c := range r.connections {
		rows[i] = table.Row{
			c.Protocol,
			c.State,
			c.RecvQ,
			c.SendQ,
			c.Local,
			c.Peer,
			c.Process,
		}
	}
	return rows
}

func (r *SSCommandRunner) Run(ctx context.Context) CommandRunResults {
	raw, err := r.command.Run(ctx)
	if err != nil {
		r.connections = []commands.Connection{}
		return err
	}

	lines := strings.Split(strings.TrimSpace(raw), "\n")
	if len(lines) < 2 {
		return nil
	}

	var connections []commands.Connection

	// Skip header line
	for _, line := range lines[1:] {
		conn, ok := parseLine(line)
		if ok {
			connections = append(connections, conn)
		}
	}
	r.connections = connections
	return nil
}

func (r *SSCommandRunner) PrintCommandAsStr() string {
	return r.command.PrintCommandAsStr()
}

func parseLine(line string) (commands.Connection, bool) {
	fields := strings.Fields(line)
	if len(fields) < 6 {
		return commands.Connection{}, false
	}

	conn := commands.Connection{
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
