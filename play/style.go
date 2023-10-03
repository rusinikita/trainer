package play

import "github.com/charmbracelet/lipgloss"

var (
	questionNumberStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Padding(0, 1)
	titleStyle          = lipgloss.NewStyle().Margin(1)
	answerF             = func() lipgloss.Style {
		return lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			BorderStyle(lipgloss.NormalBorder())
	}
	answerStyle         = answerF()
	answerSelectedStyle = answerF().
				Foreground(lipgloss.Color("170")).
				BorderStyle(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("170"))
	answerErrorStyle = answerF().
				Foreground(lipgloss.Color("160")).
				BorderStyle(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("160"))
	answerGoodStyle = answerF().
			Foreground(lipgloss.Color("70")).
			BorderForeground(lipgloss.Color("70"))
	answerGoodSelectedStyle = answerF().
				Foreground(lipgloss.Color("70")).
				BorderStyle(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("170"))
	errorStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Margin(1)
	questionTagStyle = lipgloss.NewStyle().Padding(0, 1).Margin(1).Background(lipgloss.Color("62")).Foreground(lipgloss.Color("230"))
	questionEndStyle = lipgloss.NewStyle().Margin(1).Foreground(lipgloss.Color("34"))

	listStyle         = lipgloss.NewStyle().MarginTop(1)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)
