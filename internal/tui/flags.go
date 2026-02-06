package tui

import (
	"fmt"
	"strings"

	"github.com/GabrielDCelery/netmon/internal/styles"
)

// Flag represents a command-line flag and its description.
type Flag struct {
	Name        string
	Description string
}

// SSFlags returns the flags used in the ss command.
func SSFlags() []Flag {
	return []Flag{
		{Name: "-t", Description: "Show TCP sockets"},
		{Name: "-u", Description: "Show UDP sockets"},
		{Name: "-n", Description: "Don't resolve service names"},
		{Name: "-a", Description: "Display all sockets (listening and non-listening)"},
		{Name: "-p", Description: "Show process using socket"},
	}
}

// renderFlagsPanel renders the flags panel with a border and title.
func renderFlagsPanel(flags []Flag, width, height int) string {
	var content strings.Builder

	// Add title
	content.WriteString(styles.FlagsPanelTitle.Render("Command Flags"))
	content.WriteString("\n\n")

	// Add each flag with its description
	for i, flag := range flags {
		content.WriteString(styles.FlagName.Render(flag.Name))
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
