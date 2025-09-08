package controller

import (
	"backend/internal/todo/dtos"
	"backend/internal/todo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	todo := &dtos.CreateTodoRequest{}
	if err := ctx.ShouldBindJSON(todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todoResponse, err := c.todoService.CreateTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, todoResponse)
}

func (c *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}
