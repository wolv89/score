package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	score       []LiveSet
	sets        []WonSet
	gamesPerSet uint8
}

func (m *model) ScorePoint(left bool) {

	var next LiveSet
	n := len(m.score) - 1

	if left {
		next = m.score[n].PointLeft()
	} else {
		next = m.score[n].PointRight()
	}

	// Check meta for set completion...

	m.score = append(m.score, next)

}

func (m *model) Undo() {

	n := len(m.score) - 1

	// Lots of cases to be handled, from match being finished, back to set being finished, etc

	if n == 0 {
		return
	}

	m.score = m.score[:n]

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
			m.ScorePoint(true)
		case "p":
			m.ScorePoint(false)
		case "z":
			m.Undo()
		}
	}
	return m, nil
}

func (m model) View() string {

	n := len(m.score) - 1

	return m.score[n].Render()

}

func main() {

	scoreStack := make([]LiveSet, 1)
	scoreStack[0] = NewLiveSet(6)

	m := model{
		score:       scoreStack,
		sets:        make([]WonSet, 0),
		gamesPerSet: 6,
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
