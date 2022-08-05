package styles

import "github.com/charmbracelet/lipgloss"

var (
	ParagraphTextStyle         = lipgloss.NewStyle().Margin(1, 0, 1, 4)
	ParagraphNoMarginTextStyle = lipgloss.NewStyle().Margin(0, 0, 0, 4)
	ProgressTitleTextStyle     = lipgloss.NewStyle().Margin(0, 0, 0, 4)
	QuitTextStyle              = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	SuccessTextStyle           = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	SuccessBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Margin(1, 0, 2, 4)
)
