package textinput

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/styles"
)

type Model struct {
	textInput textinput.Model
	Label     string
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.textInput.Blur()
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return styles.ParagraphNoMarginTextStyle.Render(fmt.Sprintf(
		m.Label+" %s",
		m.textInput.View(),
	) + "\n")
}

func New() Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Prompt = ""
	ti.SetCursorMode(textinput.CursorStatic)

	return Model{
		textInput: ti,
	}
}

func NewPassword() Model {
	m := New()
	m.textInput.EchoMode = textinput.EchoPassword
	return m
}
