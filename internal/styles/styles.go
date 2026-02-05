package styles

import "github.com/charmbracelet/lipgloss"

// Connection state colors.
var (
	StateEstab     = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
	StateListen    = lipgloss.NewStyle().Foreground(lipgloss.Color("#3498DB"))
	StateTimeWait  = lipgloss.NewStyle().Foreground(lipgloss.Color("#F39C12"))
	StateCloseWait = lipgloss.NewStyle().Foreground(lipgloss.Color("#E74C3C"))
	StateUnconn    = lipgloss.NewStyle().Foreground(lipgloss.Color("#95A5A6"))
	StateDefault   = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDC3C7"))
)

// UI element styles.
var (
	Title     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4")).Padding(0, 1)
	StatusBar = lipgloss.NewStyle().Foreground(lipgloss.Color("#A0A0A0")).Padding(0, 1)
	HelpText  = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
)

// StyleForState returns the appropriate lipgloss style for a connection state.
func StyleForState(state string) lipgloss.Style {
	switch state {
	case "ESTAB":
		return StateEstab
	case "LISTEN":
		return StateListen
	case "TIME-WAIT":
		return StateTimeWait
	case "CLOSE-WAIT":
		return StateCloseWait
	case "UNCONN":
		return StateUnconn
	default:
		return StateDefault
	}
}
