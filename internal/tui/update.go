package tui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

const refreshInterval = 2 * time.Second

// tickMsg signals a refresh.
type tickMsg time.Time

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
		case "?":
			m.showFlagsPanel = !m.showFlagsPanel
			m.updateTableWidth()
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.logger.Log(zap.DebugLevel, "set window dimensions", zap.Int("width", m.width), zap.Int("height", m.height))
		m.table.SetHeight(msg.Height - 4)
		m.updateTableWidth()
		m.ready = true

	case tickMsg:
		err := m.commandRunner.Run(context.Background())
		if err != nil {
			m.err = err
			return m, tick()
		}
		m.table.SetRows(m.commandRunner.Rows())
		return m, tick()
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// updateTableWidth adjusts the table width based on whether the flags panel is visible.
func (m *Model) updateTableWidth() {
	if m.showFlagsPanel && m.width > 0 {
		// Allocate 75% to table, 25% to flags panel
		// Adjust column widths proportionally
		cols := m.table.Columns()
		if len(cols) > 0 {
			// Recalculate column widths to fit the new table width
			cols[0].Width = 6  // Proto
			cols[1].Width = 10 // State
			cols[2].Width = 7  // Recv-Q
			cols[3].Width = 7  // Send-Q
			cols[4].Width = 20 // Local Address
			cols[5].Width = 20 // Peer Address
			cols[6].Width = 25 // Process
			m.table.SetColumns(cols)
		}
	} else {
		// Full width
		cols := m.table.Columns()
		if len(cols) > 0 {
			cols[0].Width = 6  // Proto
			cols[1].Width = 12 // State
			cols[2].Width = 8  // Recv-Q
			cols[3].Width = 8  // Send-Q
			cols[4].Width = 25 // Local Address
			cols[5].Width = 25 // Peer Address
			cols[6].Width = 30 // Process
			m.table.SetColumns(cols)
		}
	}
}
