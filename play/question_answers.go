package play

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"

	"github.com/rusinikita/trainer/challenge"
)

type questionModel struct {
	challenge.Question
	currentAnswer int
	foundAnswers  map[int]bool
	isErrorChoice bool

	isLast bool

	b keyBindings
}

func newQuestionModel(q challenge.Question, isLast bool) questionModel {
	return questionModel{
		Question:     q,
		foundAnswers: map[int]bool{},
		isLast:       isLast,
		b:            newBindings(),
	}
}

func (q questionModel) Update(msg tea.Msg, codeLineIndex int) questionModel {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, q.b.Left):
			q.isErrorChoice = false
			q.currentAnswer = max(0, q.currentAnswer-1)
		case key.Matches(msg, q.b.Right):
			q.isErrorChoice = false
			q.currentAnswer = min(len(q.Answers)-1, q.currentAnswer+1)

		case key.Matches(msg, q.b.Input):
			isRight := q.answer().IsRight(codeLineIndex)

			q.isErrorChoice = !isRight
			if isRight {
				q.foundAnswers[q.currentAnswer] = true
			}
		}
	}

	return q
}

func (q questionModel) View() string {
	var answers []string

	for i, answer := range q.Answers {
		style := answerStyle

		foundAnswer := q.foundAnswers[i]

		if foundAnswer {
			style = answerGoodStyle
		}

		if i == q.currentAnswer {
			style = answerSelectedStyle

			if q.isErrorChoice {
				style = answerErrorStyle
			}

			if foundAnswer {
				style = answerGoodSelectedStyle
			}
		}

		answers = append(answers, style.Render(wordwrap.String(strings.TrimSpace(answer.Text), 15)))
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render(q.Text),
		q.questionStatusLine(),
		lipgloss.JoinHorizontal(lipgloss.Left, answers...),
	)
}

func (q questionModel) questionStatusLine() string {
	if q.questionFullyAnswered() && q.isLast {
		return questionEndStyle.Render("Challenge has completed, press 'ESC' or 'q' to exit")
	} else if q.questionFullyAnswered() {
		return questionEndStyle.Render("All answers has found, press 'N' to show next question")
	}

	rightAnswers, hasLines := q.RightAnswers()
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

func (q questionModel) answer() challenge.Answer {
	return q.Answers[q.currentAnswer]
}

func (q questionModel) questionFullyAnswered() bool {
	for i, answer := range q.Answers {
		if answer.IsWrong() {
			continue
		}

		if !q.foundAnswers[i] {
			return false
		}
	}

	return true
}
