package progress

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/styles"
)

const (
	padding  = 4
	maxWidth = 80
)

type tickMsg time.Time

type Model struct {
	percent  float64
	progress progress.Model
	Title    string
}

func (Model) Init() tea.Cmd {
	return tickCmd()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		default:
			m.percent += 0.25
		}

		return m, nil

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		m.percent += 0.015
		if m.percent > 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()

	default:
		return m, nil
	}
}

func (m Model) View() string {
	pad := strings.Repeat(" ", padding)
	return styles.ProgressTitleTextStyle.Render(m.Title) + "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n"

}

func tickCmd() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewLinear() Model {
	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	m := Model{progress: prog}
	return m
}
