package infrastructure

import (
	"time"

	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/entity"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/value"
)

type TodoModel struct {
	ID      string    `bson:"_id"`
	Version int       `bson:"version"`
	Status  bool      `bson:"status"`
	Note    string    `bson:"note"`
	Due     time.Time `bson:"due"`
}

func (tm *TodoModel) Increment() {
	tm.Version++
}

type todoMapper struct {
}

func (todoMapper) FromModel(m *TodoModel) *entity.Todo {
	return entity.RestoreTodo(
		value.TodoID(m.ID),
		value.Version(m.Version),
		value.TodoStatus(m.Status),
		value.TodoNote(m.Note),
		m.Due,
	)
}

func (todoMapper) ToModel(t *entity.Todo) *TodoModel {
	return &TodoModel{
		ID:      t.ID().String(),
		Version: t.Version().Value(),
		Status:  t.Status().Value(),
		Note:    t.Note().String(),
		Due:     t.Due(),
	}
}
