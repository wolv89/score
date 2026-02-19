package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	score []LiveSet
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "o":
			m.score = append(m.score, m.score[len(m.score)-1].PointLeft())
		case "p":
			m.score = append(m.score, m.score[len(m.score)-1].PointRight())
		case "k":
			m.score = append(m.score, m.score[len(m.score)-1].GameLeft())
		case "l":
			m.score = append(m.score, m.score[len(m.score)-1].GameRight())
		case "z":
			if len(m.score) > 1 {
				m.score = m.score[:len(m.score)-1]
			}
		}
	}
	return m, nil
}

func (m model) View() string {

	return m.score[len(m.score)-1].Render()

}

func main() {
	if _, err := tea.NewProgram(model{score: make([]LiveSet, 1)}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
