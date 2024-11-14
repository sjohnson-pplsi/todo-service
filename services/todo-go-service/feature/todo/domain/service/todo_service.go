package service

import (
	"context"
	"iter"
	"time"

	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/entity"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/value"
)

type TodoService struct {
	todoRepository TodoRepository
}

type TodoRepository interface {
	GetTodo(context.Context, value.TodoID) (*entity.Todo, error)
	InsertTodo(context.Context, *entity.Todo) error
	ListTodos(context.Context, int, int) iter.Seq2[*entity.Todo, error]
	ReplaceTodo(context.Context, *entity.Todo) error
}

func NewTodoService(
	todoRepository TodoRepository,
) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

type CreateTodoCommand struct {
	Note value.TodoNote
	Due  time.Time
}

func (ts *TodoService) CreateTodo(ctx context.Context, cmd CreateTodoCommand) (value.TodoID, error) {
	todo := entity.NewTodo(value.NewTodoID(), cmd.Note, cmd.Due)
	if err := ts.todoRepository.InsertTodo(ctx, todo); err != nil {
		return "", err
	}

	return todo.ID(), nil
}

type CompleteTodoCommand struct {
	ID value.TodoID
}

func (ts *TodoService) CompleteTodo(ctx context.Context, cmd CompleteTodoCommand) error {
	todo, err := ts.todoRepository.GetTodo(ctx, cmd.ID)
	if err != nil {
		return err
	}

	todo.Complete()

	return ts.todoRepository.ReplaceTodo(ctx, todo)
}

type ResetTodoCommand struct {
	ID value.TodoID
}

func (ts *TodoService) ResetTodo(ctx context.Context, cmd ResetTodoCommand) error {
	todo, err := ts.todoRepository.GetTodo(ctx, cmd.ID)
	if err != nil {
		return err
	}

	todo.Reset()

	return ts.todoRepository.ReplaceTodo(ctx, todo)
}

type GetTodoQuery struct {
	ID value.TodoID
}

func (ts *TodoService) GetTodo(ctx context.Context, qry GetTodoQuery) (*entity.Todo, error) {
	return ts.todoRepository.GetTodo(ctx, qry.ID)
}

type ListTodosQuery struct {
	Limit  int
	Offset int
}

func (qry ListTodosQuery) FixedLimit() int {
	if qry.Limit == 0 {
		return 20
	}
	return qry.Limit
}

func (ts *TodoService) ListTodos(ctx context.Context, qry ListTodosQuery) iter.Seq2[*entity.Todo, error] {
	return ts.todoRepository.ListTodos(ctx, qry.FixedLimit(), qry.Offset)
}
