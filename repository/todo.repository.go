package repository

import (
	"github.com/techlateef/tech-lateef-gol/entities"
	"gorm.io/gorm"
)

type ToDoRepository interface {
	CreateToDo(t entities.ToDoModels) entities.ToDoModels
	UpdateTodo(t entities.ToDoModels, todoId uint64) entities.ToDoModels
	DeleteTodo(todoId uint64) entities.ToDoModels
	FindTodoById(todoId uint64) entities.ToDoModels
	GetALL() []entities.ToDoModels
}

type todoRepository struct {
	connection *gorm.DB
}

func NewToDoRepository(todoRepo *gorm.DB) ToDoRepository {
	return &todoRepository{
		connection: todoRepo,
	}

}

func (db *todoRepository) CreateToDo(t entities.ToDoModels) entities.ToDoModels {
	db.connection.Save(&t)
	db.connection.Preload("ToDoModels").Find(&t)
	db.connection.Preload("User.ToDoModels").First(&t)
	return t
}

func (db *todoRepository) UpdateTodo(t entities.ToDoModels, todoId uint64) entities.ToDoModels {
	db.connection.Save(&t)
	db.connection.Preload("ToDoModels").First(&t)

	return t
}
func (db *todoRepository) DeleteTodo(todoId uint64) entities.ToDoModels {
	var ToDoDelete entities.ToDoModels
	db.connection.Delete(&ToDoDelete, todoId)
	return ToDoDelete
}

func (db *todoRepository) GetALL() []entities.ToDoModels {
	var Todos []entities.ToDoModels
	db.connection.Preload("ToDoModels").Find(&Todos)
	return Todos
}

func (db *todoRepository) FindTodoById(todoId uint64) entities.ToDoModels {
	var todoWId entities.ToDoModels
	db.connection.Preload("ToDoModels").Find(&todoWId, todoId)
	return todoWId
}
