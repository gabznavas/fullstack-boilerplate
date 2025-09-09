package controller

import (
	"backend/internal/todo/dtos"
	"backend/internal/todo/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoUsecases usecases.TodoUsecases
}

func NewTodoController(todoUsecases usecases.TodoUsecases) *TodoController {
	return &TodoController{todoUsecases: todoUsecases}
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	todo := &dtos.CreateTodoRequest{}
	if err := ctx.ShouldBindJSON(todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todoResponse, err := c.todoUsecases.CreateTodo(ctx, todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, todoResponse)
}

func (c *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := c.todoUsecases.GetAllTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}
