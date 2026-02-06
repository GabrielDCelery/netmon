package tui

import (
	"fmt"

	"github.com/GabrielDCelery/netmon/internal/styles"
)

// View renders the TUI.
func (m Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	title := styles.Title.Render("netmon — Network Connections")
	cmdInfo := styles.CommandInfo.Render(fmt.Sprintf("Running: %s", m.runner.Command()))

	var status string
	if m.err != nil {
		status = styles.StatusBar.Render(fmt.Sprintf("Error: %v", m.err))
	} else {
		refreshTime := ""
		if !m.lastRefresh.IsZero() {
			refreshTime = fmt.Sprintf(" • Last refresh: %s", m.lastRefresh.Format("15:04:05"))
		}
		status = styles.StatusBar.Render(fmt.Sprintf("%d connections%s", len(m.connections), refreshTime))
	}

	help := styles.HelpText.Render("↑/↓: navigate • q: quit")

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s", title, cmdInfo, m.table.View(), status, help)
}
