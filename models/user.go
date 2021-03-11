package models

import (
	"errors"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:password`
}

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{Id: 1, Username: "diego", Password: "11597"}
	users[user.Id] = user
}

func GetUsers() []User {
	var list []User
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("El usuario no se encuentra dentro del mapa")
}

func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.Id] = user
	return user
}

func DeleteUser(id int) {
	delete(users, id)
}
