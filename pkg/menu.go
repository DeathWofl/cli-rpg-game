package pkg

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Menu struct {
	Choices  []string         // items on the to-do list
	Cursor   int              // which to-do list item our cursor is pointing at
	Selected map[int]struct{} // which to-do items are selected
}

func (m Menu) Init() tea.Cmd {
	return nil
}
