package repository

import (
	"github.com/techlateef/tech-lateef-gol/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entities.User) entities.User
	UpdateUser(user entities.User, userId uint64) entities.User
	GetAllUsers() []entities.User
	GetUserById(userId uint64) entities.User
	DeleteUser(user entities.User, userId uint64) entities.User
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

func (db *UserConnection) CreateUser(user entities.User) entities.User {
	db.connection.Save(&user)
	db.connection.Preload("User").Find(&user)
	return user
}

func (db *UserConnection) UpdateUser(user entities.User, userId uint64) entities.User {
	var users entities.User
	db.connection.Model(&users).Save(&user)
	db.connection.Preload("User").First(&users, userId)

	return user
}

func (db *UserConnection) GetAllUsers() []entities.User {
	var users []entities.User
	db.connection.Preload("Users").Find(&users)
	return users
}

func (db *UserConnection) GetUserById(userId uint64) entities.User {
	var user entities.User
	db.connection.Preload("user").Find(&user, userId)

	return user

}

func (db *UserConnection) DeleteUser(user entities.User, userId uint64) entities.User {
	var users entities.User
	db.connection.Preload("User").Delete(&users, userId)
	return user

}
