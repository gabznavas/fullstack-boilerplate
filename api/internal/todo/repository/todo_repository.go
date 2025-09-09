package repository

import (
	"backend/internal/todo/models"
	"context"
	"database/sql"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *models.Todo) error
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
}

type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) CreateTodo(ctx context.Context, todo *models.Todo) error {
	query := `INSERT INTO todos (id, title, completed) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.Completed)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepositoryImpl) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	query := `SELECT id, title, completed FROM todos`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	todos := []*models.Todo{}
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}
