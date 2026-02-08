package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/GabrielDCelery/netmon/internal/commands"
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
			renderFlagsPanel(m.commandRunner.GetAvailableFlags(), rightWidth, panelHeight),
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

// renderFlagsPanel renders the flags panel with a border and title.
func renderFlagsPanel(flags []commands.Flag, width, height int) string {
	var content strings.Builder

	// Add title
	content.WriteString(styles.FlagsPanelTitle.Render("Command Flags"))
	content.WriteString("\n\n")

	// Add each flag with its description
	for i, flag := range flags {
		content.WriteString(styles.FlagName.Render(flag.Short))
		content.WriteString("\n")
		content.WriteString(styles.FlagDescription.Render(fmt.Sprintf("  %s", flag.Description)))
		if i < len(flags)-1 {
			content.WriteString("\n\n")
		}
	}

	// Apply border and size constraints
	panel := styles.FlagsPanelBorder.
		Width(width - 4).
		Height(height - 4).
		Render(content.String())

	return panel
}
