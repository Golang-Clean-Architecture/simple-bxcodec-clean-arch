package http

import (
	"bxcodec-clean-arch/article/usecase"
	"bxcodec-clean-arch/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewTodoController(todoService usecase.TodoServiceImpl) TodoController {
	return TodoController{
		TodoUsecase: todoService,
	}
}

func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todo domain.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.TodoUsecase.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TodoController) GetTodo(ctx *gin.Context) {
	var taskName string = ctx.Param("name")
	todo, err := tc.TodoUsecase.GetTodo(&taskName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (tc *TodoController) GetAll(ctx *gin.Context) {
	todos, err := tc.TodoUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (tc *TodoController) UpdateTodo(ctx *gin.Context) {
	var todo domain.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.TodoUsecase.UpdateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TodoController) DeleteTodo(ctx *gin.Context) {
	taskName := ctx.Param("name")
	err := tc.TodoUsecase.DeleteTodo(&taskName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
