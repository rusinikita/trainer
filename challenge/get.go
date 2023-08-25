package challenge

import (
	"embed"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
)

//go:embed files/*
var taskFiles embed.FS

var taskValidator = customValidator()

func LoadAll() (tasks []Challenge, err error) {
	dirEntries, err := taskFiles.ReadDir("files")
	if err != nil {
		return nil, err
	}

	for _, entry := range dirEntries {
		bytes, err := taskFiles.ReadFile("files/" + entry.Name())
		if err != nil {
			return nil, err
		}

		task, err := parseTask(string(bytes))
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func customValidator() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("line_ranges", validateLineRanges)

	return validate
}

func parseTask(content string) (task Challenge, err error) {
	_, err = toml.Decode(content, &task)
	if err != nil {
		return task, err
	}

	return task, taskValidator.Struct(task)
}
