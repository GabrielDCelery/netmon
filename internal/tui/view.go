package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/GabrielDCelery/netmon/internal/styles"
)

// View renders the TUI.
func (m Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	title := styles.Title.Render("netmon — Network Connections")
	cmdInfo := styles.CommandInfo.Render(fmt.Sprintf("Running: %s", m.commandRunner.PrintCommandAsStr()))

	// Main content area - either split or full width
	var mainContent string
	panelHeight := m.height - 6 // Account for title, cmdInfo, status, help
	if m.showFlagsPanel {
		leftWidth := int(float64(m.width) * 0.75)
		rightWidth := m.width - leftWidth
		mainContent = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.table.View(),
			renderFlagsPanel(SSFlags(), rightWidth, panelHeight),
		)
	} else {
		mainContent = m.table.View()
	}

	var status string
	if m.err != nil {
		status = styles.StatusBar.Render(fmt.Sprintf("Error: %v", m.err))
	} else {
		refreshTime := ""
		if !m.lastRefresh.IsZero() {
			refreshTime = fmt.Sprintf(" • Last refresh: %s", m.lastRefresh.Format("15:04:05"))
		}
		status = styles.StatusBar.Render(fmt.Sprintf("%d connections%s", len(m.table.Rows()), refreshTime))
	}

	help := styles.HelpText.Render("↑/↓: navigate • ?: toggle flags • q: quit")

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s", title, cmdInfo, mainContent, status, help)
}
