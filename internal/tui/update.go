package tui

import (
	"context"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/GabrielDCelery/netmon/internal/netstat"
	"github.com/GabrielDCelery/netmon/internal/styles"
)

const refreshInterval = 2 * time.Second

// connectionsMsg carries the result of a connection fetch.
type connectionsMsg struct {
	connections []netstat.Connection
	err         error
}

// tickMsg signals a refresh.
type tickMsg time.Time

// fetchConnections creates a command that runs ss and parses the output.
func fetchConnections(runner netstat.Runner) tea.Cmd {
	return func() tea.Msg {
		raw, err := runner.Run(context.Background())
		if err != nil {
			return connectionsMsg{err: err}
		}
		return connectionsMsg{connections: netstat.Parse(raw)}
	}
}

// tick returns a command that sends a tickMsg after the refresh interval.
func tick() tea.Cmd {
	return tea.Tick(refreshInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles incoming messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetHeight(msg.Height - 4)
		m.ready = true

	case connectionsMsg:
		m.err = msg.err
		m.connections = msg.connections
		m.lastRefresh = time.Now()
		m.table.SetRows(connectionsToRows(msg.connections))
		return m, tick()

	case tickMsg:
		return m, fetchConnections(m.runner)
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func connectionsToRows(conns []netstat.Connection) []table.Row {
	rows := make([]table.Row, len(conns))
	for i, c := range conns {
		rows[i] = table.Row{
			c.Protocol,
			styles.StyleForState(c.State).Render(c.State),
			c.RecvQ,
			c.SendQ,
			c.Local,
			c.Peer,
			c.Process,
		}
	}
	return rows
}
