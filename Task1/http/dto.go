package http

import (
	"encoding/json"
	"errors"
	"time"
)

type completeDTO struct {
	Completed bool `json:"Completed"`
}

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) validateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type ErrorDto struct {
	Message string
	Time    time.Time
}

func (e ErrorDto) errorToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
