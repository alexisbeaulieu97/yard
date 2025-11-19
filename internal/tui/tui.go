package tui

import (
	"github.com/alexisbeaulieu97/yard/internal/domain"
	"github.com/alexisbeaulieu97/yard/internal/workspace"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	wsEngine *workspace.Engine
	tickets  []domain.Ticket
	cursor   int
	err      error
}

func NewModel(wsEngine *workspace.Engine) Model {
	return Model{
		wsEngine: wsEngine,
	}
}

func (m Model) Init() tea.Cmd {
	return m.loadTickets
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tickets)-1 {
				m.cursor++
			}
		}
	case []domain.Ticket:
		m.tickets = msg
	case error:
		m.err = msg
	}
	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	s := "Yardmaster Tickets\n\n"

	for i, t := range m.tickets {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += cursor + " " + t.ID + "\n"
	}

	s += "\nPress q to quit.\n"
	return s
}

func (m Model) loadTickets() tea.Msg {
	tickets, err := m.wsEngine.List()
	if err != nil {
		return err
	}
	return tickets
}
