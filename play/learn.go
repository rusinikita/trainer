package play

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"

	"github.com/pkg/browser"

	"github.com/rusinikita/trainer/challenge"
)

type learn struct {
	advice string
	links  []challenge.Link

	selectedLink int
	showed       bool
	b            keyBindings
}

func newLearn(advice string, links []challenge.Link) learn {
	return learn{
		advice: advice,
		links:  links,
		b:      newBindings(),
	}
}

func (l learn) Showed() bool {
	return l.showed
}

const (
	copyTitle      = "Code"
	copyAdvise     = "The code have been copied to clipboard on challenge start. Paste and play in IDE."
	challengeTitle = "Challenge"
	linksTitle     = "Useful links: ⮐  to open"
)

func (l learn) View(width int) string {
	links := make([]string, len(l.links))
	for i, link := range l.links {
		style := learningLink
		if i == l.selectedLink {
			style = learningLinkSelected
		}

		links[i] = style.Render(link.Title)
	}

	limit := width/3*2 - learningAdviseText.GetHorizontalFrameSize() - marginH1.GetHorizontalFrameSize()
	advise := lipgloss.JoinVertical(
		lipgloss.Left,
		learningTitle.Render(copyTitle),
		learningAdviseText.Render(wordwrap.String(copyAdvise, limit)),
		learningTitle.Render(challengeTitle),
		learningAdviseText.Render(wordwrap.String(l.advice, limit)),
	)
	linksStr := lipgloss.JoinVertical(
		lipgloss.Left,
		learningTitle.Render(linksTitle),
		lipgloss.JoinVertical(
			lipgloss.Left,
			links...,
		),
	)
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		marginH1.Render(advise),
		marginH1.Render(strings.Repeat("│\n", max(lipgloss.Height(advise), lipgloss.Height(linksStr)))),
		linksStr,
	)
}

func (l learn) Update(msg tea.Msg) (learn, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, l.b.Up):
			if l.selectedLink > 0 {
				l.selectedLink--
			}

		case key.Matches(msg, l.b.Down):
			if l.selectedLink < len(l.links)-1 {
				l.selectedLink++
			}

		case key.Matches(msg, l.b.Learn):
			l.showed = !l.showed
			l.selectedLink = 0

		case key.Matches(msg, l.b.Input):
			if !l.Showed() {
				return l, nil
			}

			return l, func() tea.Msg {
				_ = browser.OpenURL(l.links[l.selectedLink].URL)
				return nil
			}

		}
	}

	return l, nil
}
