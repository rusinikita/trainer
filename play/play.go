package play

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/rusinikita/trainer/challenge"
)

type model struct {
	w tea.WindowSizeMsg
	l list.Model
	b keyBindings

	challengeName   string
	questions       []challenge.Question
	currentQuestion int

	question questionModel
	learn    learn

	help help.Model
}

const copyStatus = "Code has copied to clipboard"
const copyErrStatus = "Code has not copied. Please install xsel, xclip, wl-clipboard or Termux:API"

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
	l.SetShowPagination(false)
	l.InfiniteScrolling = true
	l.StatusMessageLifetime = 5 * time.Second

	copyStatus := copyStatus

	err := clipboard.WriteAll(c.DefaultCodeSnippet)
	if err != nil {
		copyStatus = errorStyle.Render(copyErrStatus)
	}

	cmd := l.NewStatusMessage(copyStatus)

	helpModel := help.New()

	m := model{
		b: newBindings(),
		w: tea.WindowSizeMsg{
			Width:  width,
			Height: height,
		},
		l:             l,
		challengeName: c.Name,
		questions:     c.Questions,
		question:      newQuestionModel(c.Questions[0], len(c.Questions) == 1),
		learn:         newLearn(c.LearningAdvise, c.LearningLinks),
		help:          helpModel,
	}

	return m.updateListSize(), cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) allQuestionsAnswered() bool {
	return m.currentQuestion == len(m.questions)-1 && m.question.questionFullyAnswered()
}

func (m model) headerView() string {
	return ""
}

func (m model) footerView() string {
	if m.w.Height < 25 {
		return errorStyle.Render("Can't show question. Please, increase terminal window height")
	}

	number := questionNumberStyle.Render(
		fmt.Sprintf("Question %d/%d", m.currentQuestion+1, len(m.questions)),
	)

	learnSwitch := questionNumberStyle.Render("L - 👁  learn info")

	line := lipgloss.JoinHorizontal(
		lipgloss.Center,
		number,
		strings.Repeat("─", m.l.Width()-lipgloss.Width(number)-lipgloss.Width(learnSwitch)),
		learnSwitch,
	)

	var bottomView string
	if m.learn.Showed() {
		bottomView = m.learn.View(m.w.Width)
	} else {
		bottomView = m.question.View()
	}

	views := []string{
		line,
		bottomView,
		helpStyle.Render(m.help.View(m.b)),
	}

	if m.help.ShowAll || m.learn.Showed() {
		views = append(views, helpStyle.Render(copyAdvise))
	}

	return lipgloss.JoinVertical(lipgloss.Top, views...)
}

func (m model) Update(msg tea.Msg) (r tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.b.Next):
			if m.question.questionFullyAnswered() && (!m.question.isLast) {
				m.currentQuestion++
				m.question = newQuestionModel(
					m.questions[m.currentQuestion],
					m.currentQuestion == len(m.questions)-1,
				)
			}

		case key.Matches(msg, m.b.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.b.Quit):
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.w = msg
	}

	m.learn, cmd = m.learn.Update(msg)
	m.b.Learn.SetEnabled(!m.learn.Showed())
	m.b.CloseLearn.SetEnabled(m.learn.Showed())

	if cmd == nil {
		m.l, cmd = m.l.Update(msg)
	}

	if !m.learn.Showed() {
		m.question = m.question.Update(msg, m.l.Index())
	}

	m.l.Title = fmt.Sprintf("%s %d/%d", m.challengeName, m.l.Index(), len(m.l.Items()))

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
		lipgloss.Top,
		listStyle.Render(m.l.View()),
		m.footerView(),
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
