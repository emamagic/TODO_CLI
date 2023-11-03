package repository

import (
	"crypto/md5"
	"encoding/hex"
	f "fmt"
	"todo/entity"
	"todo/storage"
)

const USER_PATH = "user.txt"

var (
	userCache       []entity.User
	store                storage.Storage
	authenticatedUser entity.User
)

func init() {
	store = storage.New("")
	LoadUsers(false)
}

func LoadUsers(shouldRefresh bool) {
	if userCache == nil {
		userCache = store.LoadUserStorageFromFile()
		f.Println("Init inMemoryCache")
	}
	if shouldRefresh {
		userCache = store.LoadUserStorageFromFile()
	}
}

func AddUser(username, pass string) {

	user := entity.User{
		ID:       len(userCache) + 1,
		UserName: username,
		Pass:     hashThePass(pass),
	}

	store.SetPath(USER_PATH)
	store.WriteUserToFile(user)
	authenticatedUser = user
	LoadUsers(true)
}

func LoadUser(username, pass string) bool {

	for _, u := range userCache {
		if username == u.UserName && hashThePass(pass) == u.Pass {
			authenticatedUser = u

			return true
		}
	}
	f.Println("Username or Pass is not valid")
	return false
}

func hashThePass(pass string) string {
	hash := md5.Sum([]byte(pass))
	return hex.EncodeToString(hash[:])
}

func CurrentUser() string {
	return authenticatedUser.UserName
}

