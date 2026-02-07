package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"

	"github.com/GabrielDCelery/netmon/internal/commands"
)

type ModelOption func(*Model)

func WithLogger(logger *zap.Logger) ModelOption {
	return func(m *Model) {
		m.logger = logger
	}
}

// Model holds the application state.
type Model struct {
	table          table.Model
	commandRunner  CommandRunner
	err            error
	width          int
	height         int
	ready          bool
	lastRefresh    time.Time
	showFlagsPanel bool
	logger         *zap.Logger
}

// NewModel creates a new Model with default values.
func NewModel(opts ...ModelOption) Model {
	columns := []table.Column{
		{Title: "Proto", Width: 6},
		{Title: "State", Width: 12},
		{Title: "Recv-Q", Width: 8},
		{Title: "Send-Q", Width: 8},
		{Title: "Local Address", Width: 25},
		{Title: "Peer Address", Width: 25},
		{Title: "Process", Width: 30},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := Model{
		table:          t,
		commandRunner:  NewSSCommandRunner(commands.NewSSCommand()),
		showFlagsPanel: true,
		logger:         zap.NewNop(),
	}
	for _, opt := range opts {
		opt(&m)
	}
	return m
}

// func (m *Model) switchCommand(command CommandType) {
// 	if command == ssCommand {
// 		m.runner = commands.NewSSRunner()
// 	}
// }

// Init returns the initial command to execute.
func (m Model) Init() tea.Cmd {
	return runCommand(m.commandRunner)
}
