package service

import (
	"backend/internal/todo/dtos"
	"backend/internal/todo/models"
	"backend/internal/todo/repository"

	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(todo *dtos.CreateTodoRequest) (*dtos.CreateTodoResponse, error)
	// GetTodo(id string) (*models.Todo, error)
	GetAllTodos() ([]*models.Todo, error)
	// UpdateTodo(todo *models.Todo) error
	// DeleteTodo(id string) error
}

type TodoServiceImpl struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository}
}

func (s *TodoServiceImpl) CreateTodo(dto *dtos.CreateTodoRequest) (*dtos.CreateTodoResponse, error) {
	todo := &models.Todo{
		Title:     dto.Title,
		Completed: false,
		ID:        uuid.New().String(),
	}
	err := s.todoRepository.CreateTodo(todo)
	if err != nil {
		return nil, err
	}
	return &dtos.CreateTodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}, nil
}

func (s *TodoServiceImpl) GetAllTodos() ([]*models.Todo, error) {
	return s.todoRepository.GetAllTodos()
}
