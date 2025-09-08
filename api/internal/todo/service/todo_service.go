package service

import (
	"backend/internal/todo/cache"
	"backend/internal/todo/dtos"
	"backend/internal/todo/models"
	"backend/internal/todo/repository"
	"context"

	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(ctx context.Context, dto *dtos.CreateTodoRequest) (*dtos.CreateTodoResponse, error)
	// GetTodo(id string) (*models.Todo, error)
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
	// UpdateTodo(todo *models.Todo) error
	// DeleteTodo(id string) error
}

type TodoServiceImpl struct {
	todoRepository repository.TodoRepository
	cache          cache.TodoCache
}

func NewTodoService(todoRepository repository.TodoRepository, cache cache.TodoCache) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository, cache: cache}
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, dto *dtos.CreateTodoRequest) (*dtos.CreateTodoResponse, error) {
	todo := &models.Todo{
		Title:     dto.Title,
		Completed: false,
		ID:        uuid.New().String(),
	}
	err := s.todoRepository.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	err = s.cache.PushOneTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateTodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}, nil
}

func (s *TodoServiceImpl) GetAllTodos(ctx context.Context) (todos []*models.Todo, err error) {
	todos, err = s.cache.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	if len(todos) == 0 {
		todos, err = s.todoRepository.GetAllTodos()
		if err != nil {
			return nil, err
		}
		err = s.cache.PushAllTodos(ctx, todos)
		if err != nil {
			return nil, err
		}
	}
	return todos, nil
}
