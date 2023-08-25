package challenge

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestTomlDecoding(t *testing.T) {
	answer := Answer{}

	t.Run("nil array", func(t *testing.T) {
		_, err := toml.Decode("", &answer)
		assert.NoError(t, err)

		assert.Nil(t, answer.CodeLineRanges)
	})

	t.Run("empty array", func(t *testing.T) {
		_, err := toml.Decode("code_line_ranges=[]", &answer)
		assert.NoError(t, err)

		assert.NotNil(t, answer.CodeLineRanges)
		assert.Empty(t, answer.CodeLineRanges)
	})
}

func TestValidatorTags(t *testing.T) {
	t.Run("one of: validate enums", func(t *testing.T) {
		q := Question{
			Text: "text",
			Type: "unknown",
			Answers: []Answer{{
				Name: "name",
				Text: "text",
			}},
		}

		err := customValidator().Struct(q)

		assert.Error(t, err)
		assert.Len(t, err, 1)
	})

	t.Run("dive: validate structs in slice", func(t *testing.T) {
		q := Question{
			Text:    "text",
			Type:    "code_writing",
			Answers: []Answer{{}},
		}

		err := customValidator().Struct(q)

		assert.Error(t, err)
		assert.Len(t, err, 1)
	})

	t.Run("line_ranges", func(t *testing.T) {
		a := Answer{
			Name:           "name",
			Text:           "text",
			CodeLineRanges: [][]int{{8, 4}},
		}

		err := customValidator().Struct(a)
		assert.Error(t, err)
		assert.Len(t, err, 1)
	})
}
