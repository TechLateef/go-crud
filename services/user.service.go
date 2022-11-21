package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/techlateef/tech-lateef-gol/dto"
	"github.com/techlateef/tech-lateef-gol/entities"
	"github.com/techlateef/tech-lateef-gol/repository"
)

type UserService interface {
	CreateUser(user dto.CreateUserDto) entities.User
	UpdateUser(user dto.UpdateUserDto, userId uint64) entities.User
	GetAllUsers() []entities.User
	GetUserById(userId uint64) entities.User
	DeleteUser(user entities.User, userId uint64) entities.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserServices(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

func (service *userService) CreateUser(user dto.CreateUserDto) entities.User {
	newuser := entities.User{}
	err := smapping.FillStruct(&newuser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.userRepository.CreateUser(newuser)

	return res
}

func (service *userService) UpdateUser(user dto.UpdateUserDto, userId uint64) entities.User {
	updateuser := entities.User{}
	err := smapping.FillStruct(&updateuser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.userRepository.UpdateUser(updateuser, updateuser.Id)

	return res

}

func (service *userService) GetAllUsers() []entities.User {
	return service.userRepository.GetAllUsers()

}

func (service *userService) GetUserById(userId uint64) entities.User {

	return service.userRepository.GetUserById(userId)

}

func (service *userService) DeleteUser(user entities.User, userId uint64) entities.User {
	userToDelete := entities.User{}
	err := smapping.FillStruct(&userToDelete, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.userRepository.DeleteUser(userToDelete, userToDelete.Id)

	return res
}
