package service

import (
	"fmt"
	"todo/entity"
)

type UserRepository interface {
	GetUser() (entity.User, error)
	CreateNewUser(user entity.User) (entity.User, error)
	Login(user entity.User) error
}

type User struct {
	UserRepository UserRepository
}

func NewUserService(userRepo UserRepository) User {
	return User{
		UserRepository: userRepo,
	}
}

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	User     entity.User
	Metadata string
}

func (u User) Create(req CreateUserRequest) (CreateUserResponse, error) {
	user, error := u.UserRepository.CreateNewUser(entity.User{
		UserName: req.Username,
		PassWord: req.Password,
	})
	if error != nil {
		return CreateUserResponse{}, fmt.Errorf("can not create user: %v", error)
	}
	return CreateUserResponse{User: user}, nil
}

type LoginUserRequest struct {
	Username string
	Password string
}

func (u User) Login(req LoginUserRequest) error {
	return u.UserRepository.Login(entity.User{
		UserName: req.Username,
		PassWord: req.Password,
	})
}
