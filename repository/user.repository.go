package repository

import (
	"github.com/techlateef/tech-lateef-gol/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entities.Users) entities.Users
	UpdateUser(user entities.Users, userId uint64) entities.Users
	GetAllUsers() []entities.Users
	GetUserById(userId uint64) entities.Users
	DeleteUser(user entities.Users, userId uint64) entities.Users
}

// Creating Db instance
type UserConnection struct {
	connection *gorm.DB
}

//Creating NewUserRepository intance

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &UserConnection{
		connection: dbConn,
	}
}

func (db *UserConnection) CreateUser(user entities.Users) entities.Users {
	db.connection.Save(&user)
	db.connection.Preload("User").Find(&user)
	return user
}

func (db *UserConnection) UpdateUser(user entities.Users, userId uint64) entities.Users {
	var users entities.Users
	db.connection.Model(&users).Save(&user)
	db.connection.Preload("User").First(&users, userId)

	return user
}

func (db *UserConnection) GetAllUsers() []entities.Users {
	var users []entities.Users
	db.connection.Preload("Users").Find(&users)
	return users
}

func (db *UserConnection) GetUserById(userId uint64) entities.Users {
	var user entities.Users
	db.connection.Preload("user").Find(&user, userId)

	return user

}

func (db *UserConnection) DeleteUser(user entities.Users, userId uint64) entities.Users {
	var users entities.Users
	db.connection.Preload("User").Delete(&users, userId)
	return user

}
