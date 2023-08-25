package challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {
	tasks, err := LoadAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)

	task := tasks[0]
	assert.Equal(t, "Parallel queries", task.Name)
	assert.Equal(t, Concurrency, tasks[0].Category)
	assert.NotEmpty(t, tasks[0].Questions)
}
