package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/GabrielDCelery/netmon/internal/logger"
	"github.com/GabrielDCelery/netmon/internal/tui"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	p := tea.NewProgram(tui.NewModel(tui.WithLogger(logger)), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
