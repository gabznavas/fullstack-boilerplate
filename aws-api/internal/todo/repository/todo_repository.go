package repository

import (
	"backend/internal/todo/models"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo) error
	// GetTodo(id string) (*models.Todo, error)
	GetAllTodos() ([]*models.Todo, error)
	// UpdateTodo(todo *models.Todo) error
	// DeleteTodo(id string) error
}

type TodoRepositoryImpl struct {
	todos []*models.Todo
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{todos: []*models.Todo{}}
}

func (r *TodoRepositoryImpl) CreateTodo(todo *models.Todo) error {
	r.todos = append(r.todos, todo)
	return nil
}

// func (r *TodoRepositoryImpl) GetTodo(id string) (*models.Todo, error) {
// 	return nil, nil
// }

func (r *TodoRepositoryImpl) GetAllTodos() ([]*models.Todo, error) {
	return r.todos, nil
}

// func (r *TodoRepositoryImpl) UpdateTodo(todo *models.Todo) error {
// 	return nil
// }

// func (r *TodoRepositoryImpl) DeleteTodo(id string) error {
// 	return nil
// }
