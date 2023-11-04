package filestore

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"todo/entity"
	"errors"
)

const (
	DEFAULT_PATH = "storage.txt"
	SEPARATOR    = "&*\n"
)

type User struct {
	AuthenticatedUser entity.User
	path              string
	Users             []entity.User
}

func NewUserStore(path string) *User {
	p := DEFAULT_PATH
	if path != "" {
		p = path
	}
	user :=  &User{
		path: p,
	}
	err := user.loadUserStorageFromFile()
	if err != nil {
		fmt.Printf("can not load users from file: %v\n", err)
	}
	return user
}

func (u User) GetUser() (entity.User, error) {
	 user := u.AuthenticatedUser
	 if user.ID == 0 {
		return entity.User{}, errors.New("you should login first")
	 }
	 return user, nil
}

func (u *User) Login(us entity.User) error {
	isFound := false
	for _, user := range u.Users {
		if user.UserName == us.UserName && user.PassWord == hashThePass(us.PassWord) {
			isFound = true
			u.AuthenticatedUser = user	

			break
		}
	}
	if !isFound {
		return errors.New("username or password is not valid")
	}
	return nil
}

func (u *User) CreateNewUser(user entity.User) (entity.User, error) {
	user.ID = len(u.Users) + 1
	user.PassWord = hashThePass(user.PassWord)
	us, err := u.writeUserToFile(user)
	return us.AuthenticatedUser, err
}

func (s *User) loadUserStorageFromFile() error {

	output, rErr := os.ReadFile(s.path)
	if rErr != nil {
		return fmt.Errorf("can not read %s information %v", s.path, rErr)
	}

	data := strings.Split(strings.Trim(string(output), SEPARATOR), SEPARATOR)

	for _, u := range data {
		user := entity.User{}
		uErr := json.Unmarshal([]byte(u), &user)
		if uErr != nil {
			return fmt.Errorf("can not deserialize data %v for storing that", uErr)
		}
		s.Users = append(s.Users, user)
	}
	return nil
}

func (s *User) writeUserToFile(user entity.User) (User, error) {
	s.AuthenticatedUser = user
	var file *os.File

	file, OErr := os.OpenFile(s.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if OErr != nil {
		return User{}, fmt.Errorf("can not open the file because of %v", OErr)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("can not close file %v\n", err)

			return
		}
	}()

	data, mErr := json.Marshal(user)
	if mErr != nil {
		return User{}, fmt.Errorf("can not serialize user %v", mErr)
	}

	_, wErr := file.Write(data)
	_, wErr2 := file.WriteString(SEPARATOR)
	if wErr != nil || wErr2 != nil {
		return User{}, fmt.Errorf("can not write user to file %v", wErr)
	}

	return *s, nil
}

func hashThePass(pass string) string {
	hash := md5.Sum([]byte(pass))
	return hex.EncodeToString(hash[:])
}
