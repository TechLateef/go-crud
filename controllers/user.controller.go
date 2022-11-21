package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techlateef/tech-lateef-gol/dto"
	"github.com/techlateef/tech-lateef-gol/entities"
	"github.com/techlateef/tech-lateef-gol/services"
)

type UserController interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userServ services.UserService) UserController {
	return &userController{
		userService: userServ,
	}
}

func (user *userController) CreateUser(c *gin.Context) {
	var newUser dto.CreateUserDto
	errDto := c.ShouldBind(&newUser)
	if errDto == nil {
		c.JSON(http.StatusCreated, gin.H{
			"Message": "Empty user",
		})

	} else {
		result := user.userService.CreateUser(newUser)
		c.JSON(http.StatusCreated, result)

	}

}

func (user *userController) UpdateUser(c *gin.Context) {
	var updateuser dto.UpdateUserDto

	c.ShouldBind(&updateuser)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	updateuser.Id = id
	result := user.userService.UpdateUser(updateuser, id)

	c.JSON(http.StatusBadRequest, result)

}

func (user *userController) GetAllUsers(c *gin.Context) {
	var Users []entities.User = user.userService.GetAllUsers()

	c.JSON(http.StatusOK, Users)

}

func (user *userController) GetUserById(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Data not found",
		})
	}
	var userWithId entities.User = user.userService.GetUserById(id)
	if (userWithId == entities.User{}) {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no Data found",
		})

	} else {
		c.JSON(http.StatusOK, userWithId)
	}

}
func (user *userController) DeleteUser(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	var userId entities.User
	userId.Id = id
	res := user.userService.DeleteUser(userId, id)
	c.JSON(http.StatusBadRequest, res)

}
