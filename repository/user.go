package repository

import (
	"errors"
	"fmt"
	"todo/entity"
	"todo/storage"
)

type User struct {
	AuthenticatedUser entity.User
	Users             []entity.User
	userStorage       storage.StorageRepository
}

func NewUserRepo(userStore storage.StorageRepository) *User {
	user := User{
		Users:       make([]entity.User, 0),
		userStorage: userStore,
	}
	list, err := user.userStorage.Load()
	if err != nil {
		fmt.Printf("can not load users: %v", err)
	}
	user.Users = list

	return &user
}

func (u User) AuthUser() (entity.User, error) {
	user := u.AuthenticatedUser
	if user.ID == 0 {
		return entity.User{}, errors.New("no user logged in")
	}
	return user, nil
}

func (u *User) LoginUser(us entity.User) error {
	isFound := false
	for _, user := range u.Users {
		if user.Username == us.Username && user.Password == u.userStorage.Hash(us.Password) {
			isFound = true
			u.setAuthUser(user)

			break
		}
	}
	if !isFound {
		return errors.New("username or password is not valid")
	}
	return nil
}

func (u *User) RegisterUser(user entity.User) (entity.User, error) {
	user.ID = len(u.Users) + 1
	us, err := u.userStorage.Save(user)
	u.setAuthUser(us)
	return us, err
}

func (u *User) setAuthUser(user entity.User) {
	u.AuthenticatedUser = user
}
