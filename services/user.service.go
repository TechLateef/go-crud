package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/techlateef/tech-lateef-gol/dto"
	"github.com/techlateef/tech-lateef-gol/entities"
	"github.com/techlateef/tech-lateef-gol/repository"
)

type UserService interface {
	CreateUser(user dto.CreateUserDto) entities.Users
	UpdateUser(user dto.UpdateUserDto, userId uint64) entities.Users
	GetAllUsers() []entities.Users
	GetUserById(userId uint64) entities.Users
	DeleteUser(user entities.Users, userId uint64) entities.Users
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserServices(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

func (service *userService) CreateUser(user dto.CreateUserDto) entities.Users {
	newuser := entities.Users{}
	err := smapping.FillStruct(&newuser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.userRepository.CreateUser(newuser)

	return res
}

func (service *userService) UpdateUser(user dto.UpdateUserDto, userId uint64) entities.Users {
	updateuser := entities.Users{}
	err := smapping.FillStruct(&updateuser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.userRepository.UpdateUser(updateuser, updateuser.Id)

	return res

}

func (service *userService) GetAllUsers() []entities.Users {
	return service.userRepository.GetAllUsers()

}

func (service *userService) GetUserById(userId uint64) entities.Users {

	return service.userRepository.GetUserById(userId)

}

func (service *userService) DeleteUser(user entities.Users, userId uint64) entities.Users {
	userToDelete := entities.Users{}
	err := smapping.FillStruct(&userToDelete, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.userRepository.DeleteUser(userToDelete, userToDelete.Id)

	return res
}
