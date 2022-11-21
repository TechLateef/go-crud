package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/techlateef/tech-lateef-gol/config"
	"github.com/techlateef/tech-lateef-gol/controllers"
	"github.com/techlateef/tech-lateef-gol/repository"
	"github.com/techlateef/tech-lateef-gol/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	todoRepository repository.ToDoRepository  = repository.NewToDoRepository(db)
	service        services.ToDoService       = services.NewToDoService(todoRepository)
	userRepo       repository.UserRepository  = repository.NewUserRepository(db)
	userServ       services.UserService       = services.NewUserServices(userRepo)
	controller     controllers.TodoController = controllers.NewToDoController(service)
	control        controllers.UserController = controllers.NewUserController(userServ)
)

// Route func to server endpoints
func Routes() {

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/create", control.CreateUser)
		userRoute.GET("/all", control.GetAllUsers)
		userRoute.GET("/:id", control.GetUserById)
		userRoute.PUT("/update/:id", control.UpdateUser)
		userRoute.DELETE("/delete/:id", control.DeleteUser)

	}

	route.POST("/todo", controller.CreateToDo)
	route.GET("/todo", controller.GetALL)
	route.PUT("/todo/:id", controller.UpdateTodo)
	route.DELETE("/todo/:id", controller.DeleteTodo)
	route.GET("/todo/:id", controller.FindTodoById)

	//Run route whenever triggered
	route.Run()
}
