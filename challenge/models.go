package challenge

import (
	"github.com/go-playground/validator/v10"
)

type TaskCategory string

const (
	Beginner    TaskCategory = "beginner"
	Easy        TaskCategory = "easy"
	Concurrency TaskCategory = "concurrency"
	Performance TaskCategory = "performance"
	WTF         TaskCategory = "wtf"
)

type QuestionType string

const (
	SelectAnswers QuestionType = "select_answers"
	WriteCode     QuestionType = "code_writing"
)

type Challenge struct {
	Name               string       `toml:"name" validate:"required"`
	Category           TaskCategory `toml:"category" validate:"required,oneof=beginner easy concurrency performance wtf"`
	DefaultCodeSnippet string       `toml:"default_code_snippet" validate:"required"`
	Questions          []Question   `toml:"questions" validate:"required,dive"`
	LearningAdvise     string       `toml:"learning_advise" validate:"required"`
	LearningLinks      []Link       `toml:"learning_links" validate:"dive"`
}

type Question struct {
	Text    string       `toml:"text" validate:"required"`
	Type    QuestionType `toml:"type" validate:"required,oneof=select_answers code_writing"`
	Answers []Answer     `toml:"answers" validate:"required,dive"`
}

func (q Question) RightAnswers() (count int, hasLines bool) {
	for _, answer := range q.Answers {
		if answer.CodeLineRanges == nil {
			continue
		}

		count++

		if len(answer.CodeLineRanges) > 0 {
			hasLines = true
		}
	}

	return count, hasLines
}

type Answer struct {
	Name           string     `toml:"name"`
	Text           string     `toml:"text" validate:"required"`
	CodeLineRanges LineRanges `toml:"code_line_ranges" validate:"line_ranges"`
}

func (a Answer) IsRight(line int) bool {
	return a.CodeLineRanges.In(line)
}

func (a Answer) IsWrong() bool {
	return a.CodeLineRanges == nil
}

type LineRanges [][]int

func (r LineRanges) In(line int) bool {
	if r == nil {
		return false
	}

	if len(r) == 0 {
		return true
	}

	for _, lines := range r {
		first := lines[0]
		last := first

		if len(lines) == 2 {
			last = lines[1]
		}

		if line >= first && line <= last {
			return true
		}
	}

	return false
}

func (r LineRanges) validate() bool {
	if len(r) == 0 {
		return true
	}

	for _, lines := range r {
		length := len(lines)
		if length == 1 {
			continue
		}

		if length != 2 {
			return false // line range must contain 1 or 2 line numbers
		}

		if lines[0] > lines[1] {
			return false // end line must be greater than start
		}
	}

	return true
}

func validateLineRanges(fl validator.FieldLevel) bool {
	ranges, ok := fl.Field().Interface().(LineRanges)
	if !ok {
		return false
	}

	return ranges.validate()
}

type Link struct {
	Title string `toml:"title" validate:"required"`
	URL   string `toml:"url" validate:"required"`
}
