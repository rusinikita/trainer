package choose

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/rusinikita/trainer/challenge"
	"github.com/rusinikita/trainer/play"
)

type choose struct {
	err  error
	list list.Model
}

func New() tea.Model {
	listModel := list.New(nil, itemDelegate{}, 20, 14)
	listModel.SetFilteringEnabled(false)
	listModel.SetShowStatusBar(false)
	listModel.StatusMessageLifetime = 5 * time.Second
	listModel.Title = "Select challenge"
	listModel.InfiniteScrolling = true

	return choose{
		list: listModel,
	}
}

func (c choose) Init() tea.Cmd {
	return tea.Batch(c.list.StartSpinner(), func() tea.Msg {
		tasks, err := challenge.LoadAll()
		if err != nil {
			return err
		}

		return tasks
	})
}

func (c choose) Update(msg tea.Msg) (m tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		c.err = msg

	case []challenge.Challenge:
		items := make([]list.Item, 0, len(msg))
		for _, task := range msg {
			items = append(items, item(task))
		}

		c.list.StopSpinner()
		c.list.SetItems(items)

	case tea.WindowSizeMsg:
		c.list.SetSize(
			msg.Width-listStyle.GetHorizontalFrameSize(),
			msg.Height-listStyle.GetVerticalFrameSize(),
		)

	case tea.KeyMsg:
		if play.ValidateBindings(msg) {
			return c, c.list.NewStatusMessage(play.RenderLayoutErr(msg))
		}
		if msg.String() == "enter" {
			return play.New(
				challenge.Challenge(c.list.SelectedItem().(item)),
				c.list.Width()+listStyle.GetHorizontalFrameSize(),
				c.list.Height()+listStyle.GetVerticalFrameSize(),
				func(msg tea.WindowSizeMsg) (m tea.Model, cmd tea.Cmd) {
					return c.Update(msg)
				},
			)
		}
	}

	c.list, cmd = c.list.Update(msg)

	return c, cmd
}

func (c choose) View() string {
	if c.err != nil {
		return c.err.Error()
	}

	return listStyle.Render(c.list.View())
}

type item challenge.Challenge

func (i item) FilterValue() string {
	return i.Name
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}

var (
	listStyle         = lipgloss.NewStyle().Margin(1, 2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)
