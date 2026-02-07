package tui

import (
	"context"

	"github.com/GabrielDCelery/netmon/internal/commands"
	"github.com/charmbracelet/bubbles/table"
)

type SSRunner struct {
	command     *commands.SSCommand
	connections []commands.Connection
}

func NewSSRunner() *SSRunner {
	return &SSRunner{
		command:     commands.NewSSCommand(),
		connections: []commands.Connection{},
	}
}

func (r *SSRunner) Run(ctx context.Context) error {
	connections, err := r.command.Run(ctx)
	if err != nil {
		r.connections = []commands.Connection{}
		return err
	}
	r.connections = connections
	return nil
}

func (r *SSRunner) Columns() []table.Column {
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

func (r *SSRunner) Rows() []table.Row {
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

func (r *SSRunner) PrintCommandAsStr() string {
	return r.command.PrintCommandAsStr()
}
