package services

import (
	"log"

	"github.com/mashingan/smapping"

	"github.com/techlateef/tech-lateef-gol/dto"
	"github.com/techlateef/tech-lateef-gol/entities"

	"github.com/techlateef/tech-lateef-gol/repository"
)

type ToDoService interface {
	CreateToDo(t dto.CreateToDoDto) entities.ToDoModels
	UpdateTodo(t dto.UpdateToDoDto, todoId uint64) entities.ToDoModels
	DeleteTodo(todoId uint64) entities.ToDoModels
	FindTodoById(todoId uint64) entities.ToDoModels
	GetALL() []entities.ToDoModels
}

type todoService struct {
	todoRepository repository.ToDoRepository
}

func NewToDoService(todoServ repository.ToDoRepository) ToDoService {
	return &todoService{
		todoRepository: todoServ,
	}
}

func (service *todoService) CreateToDo(t dto.CreateToDoDto) entities.ToDoModels {
	TodoC := entities.ToDoModels{}
	err := smapping.FillStruct(&TodoC, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.todoRepository.CreateToDo(TodoC)

	return res
}

func (service *todoService) UpdateTodo(t dto.UpdateToDoDto, todoId uint64) entities.ToDoModels {
	todoU := entities.ToDoModels{}
	err := smapping.FillStruct(&todoU, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	return todoU
}

func (service *todoService) DeleteTodo(todoId uint64) entities.ToDoModels {
	return service.todoRepository.DeleteTodo(todoId)

}
func (service *todoService) GetALL() []entities.ToDoModels {

	return service.todoRepository.GetALL()
}

func (service *todoService) FindTodoById(todoId uint64) entities.ToDoModels {
	return service.todoRepository.FindTodoById(todoId)
}
