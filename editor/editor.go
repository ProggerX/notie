package editor

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Term      string
	Width     int
	Height    int
	NameInput textinput.Model
	TextArea  textarea.Model
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.NameInput.Focused() && !isFileExist(m.NameInput.Value()) {
				m.NameInput.Blur()
				m.TextArea.Focus()
				return m, nil
			} else if m.NameInput.Focused() {
				m.NameInput.Placeholder = "This file already exist! Choose other name, please!"
				m.NameInput.SetValue("")
			}
		case tea.KeyCtrlC:
			if m.NameInput.Value() != "" && m.TextArea.Value() != "" {
				_ = os.WriteFile(".notie/notes/"+m.NameInput.Value(), []byte(m.TextArea.Value()), os.ModePerm)
				return m, tea.Quit
			}
		case tea.KeyCtrlX:
			return m, tea.Quit
		}
	}

	m.NameInput, cmd = m.NameInput.Update(msg)
	m.TextArea, cmd = m.TextArea.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Enter name of note here:\n\n%s\n\nAnd a note here:\n\n%s\n\nWhen you're done, press ctrl+c\nOr, if you don't want to write or save anything, press ctrl+x\n\n",
		m.NameInput.View(),
		m.TextArea.View(),
	)
}
