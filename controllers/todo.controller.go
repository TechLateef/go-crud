package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techlateef/tech-lateef-gol/dto"
	"github.com/techlateef/tech-lateef-gol/entities"
	"github.com/techlateef/tech-lateef-gol/services"
)

type TodoController interface {
	CreateToDo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
	FindTodoById(c *gin.Context)
	GetALL(c *gin.Context)
}

type todoController struct {
	todoService services.ToDoService
}

func NewToDoController(todoCon services.ToDoService) TodoController {
	return &todoController{
		todoService: todoCon,
	}
}

func (todo *todoController) CreateToDo(c *gin.Context) {
	todoC := dto.CreateToDoDto{}
	errDto := c.ShouldBind(&todoC)
	if errDto != nil {
		c.JSON(http.StatusBadRequest, errDto)
	}

	res := todo.todoService.CreateToDo(todoC)
	c.JSON(http.StatusOK, res)

}

func (todo *todoController) UpdateTodo(c *gin.Context) {
	todoU := dto.UpdateToDoDto{}
	c.ShouldBind(&todoU)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	todoU.UserID = id
	res := todo.todoService.UpdateTodo(todoU, id)
	c.JSON(http.StatusOK, res)

}

func (todo *todoController) GetALL(c *gin.Context) {
	var todos []entities.ToDoModels = todo.todoService.GetALL()

	c.JSON(http.StatusOK, todos)

}

func (todo *todoController) FindTodoById(c *gin.Context) {
	var todos entities.ToDoModels
	c.ShouldBind(&todos)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	todos.ID = id
	res := todo.todoService.FindTodoById(id)
	c.JSON(http.StatusOK, res)

}

func (todo *todoController) DeleteTodo(c *gin.Context) {
	var dotos entities.ToDoModels
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	dotos.ID = id
	res := todo.todoService.DeleteTodo(id)
	c.JSON(http.StatusOK, res)
}
