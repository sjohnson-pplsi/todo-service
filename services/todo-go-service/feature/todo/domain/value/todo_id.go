package value

import "github.com/google/uuid"

type TodoID string

func (t TodoID) String() string { return string(t) }

func NewTodoID() TodoID { return TodoID(uuid.NewString()) }
