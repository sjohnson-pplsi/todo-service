package controller

import (
	"context"

	pb "github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib/generated"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

// CompleteTodo implements togo_go_lib.TodoServiceServer.
func (t *TodoService) CompleteTodo(context.Context, *pb.CompleteTodoRequest) (*pb.CompleteTodoResponse, error) {
	panic("unimplemented")
}

// CreateTodo implements togo_go_lib.TodoServiceServer.
func (t *TodoService) CreateTodo(context.Context, *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	panic("unimplemented")
}

// GetTodo implements togo_go_lib.TodoServiceServer.
func (t *TodoService) GetTodo(context.Context, *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	panic("unimplemented")
}

// ListTodos implements togo_go_lib.TodoServiceServer.
func (t *TodoService) ListTodos(context.Context, *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	panic("unimplemented")
}

// ResetTodo implements togo_go_lib.TodoServiceServer.
func (t *TodoService) ResetTodo(context.Context, *pb.ResetTodoRequest) (*pb.ResetTodoResponse, error) {
	panic("unimplemented")
}
