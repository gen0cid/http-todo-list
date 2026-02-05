package todo

import "time"

type Task struct {
	Title       string `json:"title"`
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title, decription string) Task {
	return Task{
		Title:       title,
		Description: decription,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completeTime
}

func (t *Task) UnComplete() {
	t.Completed = false
	t.CompletedAt = nil
}
