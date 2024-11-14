package entity

import (
	"time"

	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/value"
)

type Todo struct {
	id      value.TodoID
	version value.Version
	status  value.TodoStatus
	note    value.TodoNote
	due     time.Time
}

func (t *Todo) ID() value.TodoID         { return t.id }
func (t *Todo) Version() value.Version   { return t.version }
func (t *Todo) Status() value.TodoStatus { return t.status }
func (t *Todo) Note() value.TodoNote     { return t.note }
func (t *Todo) Due() time.Time           { return t.due }

func RestoreTodo(
	id value.TodoID,
	version value.Version,
	status value.TodoStatus,
	note value.TodoNote,
	due time.Time,
) *Todo {
	return &Todo{
		id:      id,
		version: version,
		status:  status,
		note:    note,
		due:     due,
	}
}

func NewTodo(
	id value.TodoID,
	note value.TodoNote,
	due time.Time,
) *Todo {
	return &Todo{
		id:      id,
		version: 0,
		status:  value.TodoStatusIncomplete,
		note:    note,
		due:     due,
	}
}

func (t *Todo) Complete() {
	t.status = value.TodoStatusComplete
}

func (t *Todo) Reset() {
	t.status = value.TodoStatusIncomplete
}

func (t *Todo) ChangeNote(note value.TodoNote) {
	t.note = note
}

func (t *Todo) ChangeDue(due time.Time) {
	t.due = due
}
