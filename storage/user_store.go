package storage

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"todo/entity"
)

const (
	DEFAULT_PATH = "storage.txt"
	SEPARATOR    = "&*\n"
)

type StorageRepository interface {
	Save(userInput entity.User) (entity.User, error)
	Load() ([]entity.User, error)
	Hash(password string) string
}

type User struct {
	path string
}

func New(path string) *User {
	p := DEFAULT_PATH
	if path != "" {
		p = path
	}
	user := &User{
		path: p,
	}
	return user
}

func (u User) Save(userInput entity.User) (entity.User, error) {
	var file *os.File

	file, OErr := os.OpenFile(u.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if OErr != nil {
		return entity.User{}, fmt.Errorf("can not open the file because of %v", OErr)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("can not close file %v\n", err)

			return
		}
	}()
	userInput.PassWord = u.Hash(userInput.PassWord)
	userOut := userInput
	data, mErr := json.Marshal(userOut)
	if mErr != nil {
		return entity.User{}, fmt.Errorf("can not serialize user %v", mErr)
	}

	_, wErr := file.Write(data)
	_, wErr2 := file.WriteString(SEPARATOR)
	if wErr != nil || wErr2 != nil {
		return entity.User{}, fmt.Errorf("can not write user to file %v", wErr)
	}

	return userOut, nil
}

func (u User) Load() ([]entity.User, error) {
	var users []entity.User
	if _, err := os.Stat(u.path); err != nil {
		return users, nil
	}

	output, rErr := os.ReadFile(u.path)
	if rErr != nil {
		return nil, fmt.Errorf("can not read %s information %v", u.path, rErr)
	}

	data := strings.Split(strings.Trim(string(output), SEPARATOR), SEPARATOR)

	for _, u := range data {
		user := entity.User{}
		uErr := json.Unmarshal([]byte(u), &user)
		users = append(users, user)
		if uErr != nil {
			return nil, fmt.Errorf("can not deserialize data %v for storing that", uErr)
		}
	}
	return users, nil
}

func (u User) Hash(pass string) string {
	hash := md5.Sum([]byte(pass))
	return hex.EncodeToString(hash[:])
}
