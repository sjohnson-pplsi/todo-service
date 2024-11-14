package infrastructure

import (
	"context"
	"iter"
	"time"

	"github.com/cardboardrobots/esmongo"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/entity"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/service"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/value"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoModel struct {
	ID      string    `bson:"id"`
	Version int       `bson:"version"`
	Status  bool      `bson:"status"`
	Note    string    `bson:"note"`
	Due     time.Time `bson:"due"`
}

type TodoRepository struct {
	repository *esmongo.MongoRepository[entity.Todo, value.TodoID, TodoModel]
}

var _ service.TodoRepository = (*TodoRepository)(nil)

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

func NewTodoRepository(collection *mongo.Collection) *TodoRepository {
	return &TodoRepository{
		repository: esmongo.NewMongoRepository[entity.Todo, value.TodoID, TodoModel](collection, todoMapper{}),
	}
}

func (t *TodoRepository) GetTodo(ctx context.Context, id value.TodoID) (*entity.Todo, error) {
	return t.repository.Get(ctx, id)
}

func (t *TodoRepository) InsertTodo(ctx context.Context, todo *entity.Todo) error {
	return t.repository.Insert(ctx, todo.ID(), todo)
}

func (t *TodoRepository) ListTodos(ctx context.Context, limit int, offset int) iter.Seq2[*entity.Todo, error] {
	return t.repository.GetList(ctx,
		esmongo.Filter(),
		esmongo.Sort().
			SetLimit(limit).
			SetSkip(offset))
}

func (t *TodoRepository) ReplaceTodo(ctx context.Context, todo *entity.Todo) error {
	return t.repository.Replace(ctx, todo.ID(), todo.Version().Value(), todo)
}
