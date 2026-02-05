package tui

import (
	"fmt"

	"github.com/GabrielDCelery/diagnoose/internal/styles"
)

// View renders the TUI.
func (m Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	title := styles.Title.Render("diagnoose — Network Connections")

	var status string
	if m.err != nil {
		status = styles.StatusBar.Render(fmt.Sprintf("Error: %v", m.err))
	} else {
		status = styles.StatusBar.Render(fmt.Sprintf("%d connections", len(m.connections)))
	}

	help := styles.HelpText.Render("↑/↓: navigate • q: quit")

	return fmt.Sprintf("%s\n%s\n%s\n%s", title, m.table.View(), status, help)
}
