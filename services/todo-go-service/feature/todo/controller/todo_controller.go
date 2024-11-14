package controller

import (
	"context"

	pb "github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib/generated"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/entity"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/service"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/value"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TodoController struct {
	pb.UnimplementedTodoServiceServer
	todoService *service.TodoService
}

func NewTodoController(
	todoService *service.TodoService,
) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (t *TodoController) CompleteTodo(ctx context.Context, request *pb.CompleteTodoRequest) (*pb.CompleteTodoResponse, error) {
	err := t.todoService.CompleteTodo(ctx, service.CompleteTodoCommand{})
	if err != nil {
		return nil, err
	}

	return &pb.CompleteTodoResponse{}, nil
}

func (t *TodoController) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	id, err := t.todoService.CreateTodo(ctx, service.CreateTodoCommand{
		Note: value.TodoNote(request.Note),
		Due:  request.Due.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTodoResponse{
		TodoId: id.String(),
	}, nil
}

// GetTodo implements togo_go_lib.TodoServiceServer.
func (t *TodoController) GetTodo(ctx context.Context, request *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	todo, err := t.todoService.GetTodo(ctx, service.GetTodoQuery{
		ID: value.TodoID(request.TodoId),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetTodoResponse{
		Todo: todoToDto(todo),
	}, nil
}

func (t *TodoController) ListTodos(ctx context.Context, request *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	data := make([]*pb.Todo, 0)
	for todo, err := range t.todoService.ListTodos(ctx, service.ListTodosQuery{}) {
		if err != nil {
			return nil, err
		}

		data = append(data, todoToDto(todo))
	}

	return &pb.ListTodosResponse{
		Count: int32(len(data)),
		Data:  data,
	}, nil
}

func (t *TodoController) ResetTodo(ctx context.Context, request *pb.ResetTodoRequest) (*pb.ResetTodoResponse, error) {
	err := t.todoService.ResetTodo(ctx, service.ResetTodoCommand{})
	if err != nil {
		return nil, err
	}

	return &pb.ResetTodoResponse{}, nil
}

func todoToDto(todo *entity.Todo) *pb.Todo {
	return &pb.Todo{
		TodoId:  todo.ID().String(),
		Version: int32(todo.Version()),
		Note:    todo.Note().String(),
		Due:     timestamppb.New(todo.Due()),
		Status:  todoStatusToDto(todo.Status()),
	}
}

func todoStatusToDto(status value.TodoStatus) pb.TodoStatus {
	switch status {
	case value.TodoStatusComplete:
		return pb.TodoStatus_TODO_STATUS_COMPLETE
	case value.TodoStatusIncomplete:
		fallthrough
	default:
		return pb.TodoStatus_TODO_STATUS_INCOMPLETE
	}
}
