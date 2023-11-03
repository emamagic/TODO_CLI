package storage

import (
	"encoding/json"
	f "fmt"
	"os"
	"strings"
	"todo/entity"
)

type Storage struct {
	path string
}

const (
	DEFAULT_PATH = "storage.txt"
	SEPARATOR    = "&*\n"
)
 
var len int

func New(path string) Storage {
	s := Storage{}
	s.SetPath(path)
	return s
}

func (s *Storage) SetPath(path string) {
	if path == "" {
		s.path = DEFAULT_PATH
	}
	s.path = path
}

func (s *Storage) Len() int {
	return len
}

func (s *Storage) LoadUserStorageFromFile() []entity.User {
	var users []entity.User

	output, rErr := os.ReadFile(s.path)
	if rErr != nil {
		f.Printf("Can not read %s information %v", s.path, rErr)

		return nil
	}

	data := strings.Split(strings.Trim(string(output), SEPARATOR), SEPARATOR)

	for _, u := range data {
		len++
		user := entity.User{}
		uErr := json.Unmarshal([]byte(u), &user)
		if uErr != nil {
			f.Printf("Can not deserialize data %v\n for storing that", uErr)

			return nil
		}
		users = append(users, user)
	}

	return users
}

func (s *Storage) WriteUserToFile(user entity.User) {
	var file *os.File

	file, OErr := os.OpenFile(s.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if OErr != nil {
		f.Printf("Can not open the file because of %v\n", OErr)

		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			f.Printf("Can not close file %v\n", err)

			return
		}
	}()

	data, mErr := json.Marshal(user)
	if mErr != nil {
		f.Printf("Can not serialize user %v\n", mErr)

		return
	}

	_, wErr := file.Write(data)
	_, wErr2 := file.WriteString(SEPARATOR)
	if wErr != nil || wErr2 != nil {
		f.Printf("can not write user to file %v\n", wErr)

		return
	}

}
