package play

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"

	"github.com/rusinikita/trainer/challenge"
)

type model struct {
	w tea.WindowSizeMsg
	l list.Model
	b keyBindings

	questions       []challenge.Question
	currentQuestion int
	currentAnswer   int
	foundAnswers    map[int]bool
	isErrorChoice   bool
}

func New(c challenge.Challenge, width, height int) (tea.Model, tea.Cmd) {
	// normalize line endings
	codeLines := lines(c.DefaultCodeSnippet)

	l := list.New(codeLines, itemDelegate{}, width, height)
	l.Title = c.Name
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.KeyMap.NextPage = key.NewBinding(key.WithDisabled())
	l.KeyMap.PrevPage = key.NewBinding(key.WithDisabled())
	l.SetShowHelp(false)
	l.InfiniteScrolling = true
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return newBindings().ShortHelp()
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return newBindings().FullHelp()[0]
	}
	l.Styles.HelpStyle = l.Styles.HelpStyle.MarginBottom(1)
	l.StatusMessageLifetime = 5 * time.Second

	copyStatus := "Code is copied to clipboard"

	err := clipboard.WriteAll(c.DefaultCodeSnippet)
	if err != nil {
		copyStatus = errorStyle.Render("Code isn't copied. Please install xsel, xclip, wl-clipboard or Termux:API.")
	}

	cmd := l.NewStatusMessage(copyStatus)

	m := model{
		b: newBindings(),
		w: tea.WindowSizeMsg{
			Width:  width,
			Height: height,
		},
		l:            l,
		questions:    c.Questions,
		foundAnswers: map[int]bool{},
	}

	return m.updateListSize(), cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) question() challenge.Question {
	return m.questions[m.currentQuestion]
}

func (m model) answer() challenge.Answer {
	return m.question().Answers[m.currentAnswer]
}

func (m model) questionFullyAnswered() bool {
	for i, answer := range m.question().Answers {
		if answer.IsWrong() {
			continue
		}

		if !m.foundAnswers[i] {
			return false
		}
	}

	return true
}

func (m model) allQuestionsAnswered() bool {
	return m.currentQuestion == len(m.questions)-1 && m.questionFullyAnswered()
}

func (m model) headerView() string {
	return ""
}

func (m model) footerView() string {
	if m.w.Height < 25 {
		return errorStyle.Render("Can't show question. Please, increase terminal window height")
	}

	question := m.question()

	number := questionNumberStyle.Render(
		fmt.Sprintf("Question %d/%d", m.currentQuestion+1, len(m.questions)),
	)

	line := lipgloss.JoinHorizontal(
		lipgloss.Center,
		number,
		strings.Repeat("─", m.l.Width()-lipgloss.Width(number)),
	)

	var answers []string

	for i, answer := range question.Answers {
		style := answerStyle

		foundAnswer := m.foundAnswers[i]

		if foundAnswer {
			style = answerGoodStyle
		}

		if i == m.currentAnswer {
			style = answerSelectedStyle

			if m.isErrorChoice {
				style = answerErrorStyle
			}

			if foundAnswer {
				style = answerGoodSelectedStyle
			}
		}

		answers = append(answers, style.Render(wordwrap.String(strings.TrimSpace(answer.Text), 15)))
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		line,
		titleStyle.Render(question.Text),
		m.questionStatusLine(),
		lipgloss.JoinHorizontal(lipgloss.Left, answers...),
		m.helpView(),
	)
}

func (m model) helpView() string {
	return m.l.Styles.HelpStyle.Render(m.l.Help.View(m.l))
}

func (m model) Update(msg tea.Msg) (r tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.b.Next):
			if m.questionFullyAnswered() {
				m.currentQuestion++
				m.currentAnswer = 0
				m.foundAnswers = map[int]bool{}
			}

		case key.Matches(msg, m.b.Left):
			m.isErrorChoice = false
			m.currentAnswer = max(0, m.currentAnswer-1)
		case key.Matches(msg, m.b.Right):
			m.isErrorChoice = false
			m.currentAnswer = min(len(m.question().Answers)-1, m.currentAnswer+1)

		case key.Matches(msg, m.b.Input):
			isRight := m.answer().IsRight(m.l.Index())

			m.isErrorChoice = !isRight
			if isRight {
				m.foundAnswers[m.currentAnswer] = true
			}
		}

	case tea.WindowSizeMsg:
		m.w = msg
	}

	m.l, cmd = m.l.Update(msg)

	return m.updateListSize(), cmd
}

func (m model) updateListSize() model {
	footerHeight := lipgloss.Height(m.footerView())
	listMargin := listStyle.GetVerticalFrameSize()

	verticalMarginHeight := footerHeight + listMargin

	m.l.SetSize(m.w.Width-listStyle.GetHorizontalFrameSize(), m.w.Height-verticalMarginHeight)

	return m
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		listStyle.Render(m.l.View()),
		m.footerView(),
	)
}

func (m model) questionStatusLine() string {
	if m.allQuestionsAnswered() {
		return questionEndStyle.Render("Challenge has completed, press 'ESC' or 'q' to exit")
	} else if m.questionFullyAnswered() {
		return questionEndStyle.Render("All answers has found, press 'N' to show next question")
	}

	rightAnswers, hasLines := m.question().RightAnswers()
	if rightAnswers == 0 {
		return ""
	}

	countTag := "one answer"
	if rightAnswers > 1 {
		countTag = "few answers"
	}

	linesTag := "with any line"
	if hasLines {
		linesTag = "with related line selected"
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		questionTagStyle.Render(countTag),
		questionTagStyle.Render(linesTag),
	)
}

type codeLine string

func lines(s string) []list.Item {
	ss := strings.Split(strings.ReplaceAll(strings.TrimSpace(s), "\r\n", "\n"), "\n")

	l := make([]list.Item, 0, len(ss))
	for _, s := range ss {
		l = append(l, codeLine(strings.ReplaceAll(s, "\t", "    ")))
	}

	return l
}

func (l codeLine) FilterValue() string {
	return string(l)
}

type keyBindings struct {
	Input key.Binding
	Left  key.Binding
	Right key.Binding
	Next  key.Binding
}

func (k keyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right, k.Input}
}

func (k keyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Left, k.Right, k.Input}}
}

func newBindings() keyBindings {
	return keyBindings{
		Left:  key.NewBinding(key.WithKeys("left"), key.WithHelp("←", "left")),
		Right: key.NewBinding(key.WithKeys("right"), key.WithHelp("→", "right")),
		Input: key.NewBinding(key.WithKeys("enter", " "), key.WithHelp("⮐", "select")),
		Next:  key.NewBinding(key.WithKeys("n"), key.WithHelp("N", "next")),
	}
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(codeLine)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index, i)

	if index < 10 {
		str = " " + str
	}

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}

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
	questionTagStyle = lipgloss.NewStyle().Padding(0, 1).Margin(1).Background(lipgloss.Color("62"))
	questionEndStyle = lipgloss.NewStyle().Margin(1).Foreground(lipgloss.Color("34"))
)

var (
	listStyle         = lipgloss.NewStyle().MarginTop(1)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)
