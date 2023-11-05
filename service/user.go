package service

import (
	"fmt"
	"todo/entity"
)

type UserRepository interface {
	AuthUser() (entity.User, error)
	RegisterUser(user entity.User) (entity.User, error)
	LoginUser(user entity.User) error
}

type User struct {
	UserRepository UserRepository
}

func NewUserService(userRepo UserRepository) User {
	return User{
		UserRepository: userRepo,
	}
}

type RegisterUserRequest struct {
	Username string
	Password string
}

type RegisterUserResponse struct {
	User     entity.User
	Metadata string
}

func (u User) Register(req RegisterUserRequest) (RegisterUserResponse, error) {
	user, error := u.UserRepository.RegisterUser(entity.User{
		UserName: req.Username,
		PassWord: req.Password,
	})
	if error != nil {
		return RegisterUserResponse{}, fmt.Errorf("can not create user: %v", error)
	}
	return RegisterUserResponse{User: user}, nil
}

type LoginUserRequest struct {
	Username string
	Password string
}

func (u User) Login(req LoginUserRequest) error {
	return u.UserRepository.LoginUser(entity.User{
		UserName: req.Username,
		PassWord: req.Password,
	})
}
