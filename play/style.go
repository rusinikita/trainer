package play

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	questionNumberStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Padding(0, 1)
	titleStyle          = lipgloss.NewStyle().Margin(1)
	answerStyle         = lipgloss.NewStyle().
				PaddingLeft(1).PaddingRight(1).
				BorderStyle(lipgloss.NormalBorder())
	warningMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("231")).
				Background(lipgloss.Color("166")).
				Padding(0, 1)
	answerSelectedStyle = answerStyle.Copy().
				Foreground(lipgloss.Color("170")).
				BorderStyle(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("170"))
	answerErrorStyle = answerStyle.Copy().
				Foreground(lipgloss.Color("160")).
				BorderStyle(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("160"))
	answerGoodStyle = answerStyle.Copy().
			Foreground(lipgloss.Color("70")).
			BorderForeground(lipgloss.Color("70"))
	answerGoodSelectedStyle = answerStyle.Copy().
				Foreground(lipgloss.Color("70")).
				BorderStyle(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("170"))
	errorStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Margin(1)
	tagStyle         = lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("62")).Foreground(lipgloss.Color("230"))
	questionTagStyle = tagStyle.Copy().Margin(1)
	questionEndStyle = lipgloss.NewStyle().Margin(1).Foreground(lipgloss.Color("34"))

	listStyle         = lipgloss.NewStyle().MarginTop(1)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

	learningTitle      = lipgloss.NewStyle().Bold(true).Italic(true)
	learningAdviseText = lipgloss.NewStyle().Padding(1, 0, 1, 4)
	marginH1           = lipgloss.NewStyle().Margin(0, 1)
	helpStyle          = lipgloss.NewStyle().Margin(1).MarginBottom(0).
				Foreground(lipgloss.AdaptiveColor{Light: "#B2B2B2", Dark: "#4A4A4A"})
	learningLink         = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	learningLinkSelected = learningLink.Copy().
				Foreground(lipgloss.Color("170")).
				BorderForeground(lipgloss.Color("170"))
)

func RenderLayoutErr(msg tea.KeyMsg) string {
	return warningMessageStyle.Render(
		fmt.Sprintf(
			"'%s' have been pressed. Please, change keyboard layout to english.",
			msg.String(),
		),
	)
}
