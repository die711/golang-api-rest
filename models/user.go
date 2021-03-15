package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type User struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:password`
	Email       string `json:email`
	createdDate time.Time
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const UserSchema string = `create table users(
	id int(6) auto_increment primary key,
    username varchar(30) not null unique,
    password varchar(64) not null,
    email varchar(40),
    created_date timestamp default current_timestamp)`

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	user.SetPassword(password)
	return user
}

func (u *User) SetPassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(hash)

}

func CreateUser(username, password, email string) (*User, error) {
	user := NewUser(username, password, email)
	err := user.Save()
	return user, err
}

func (u *User) Save() error {
	if u.Id == 0 {
		return u.insert()
	} else {
		return u.update()
	}
}

func (u *User) insert() error {
	sql := "Insert users set username=?, password=?, email=?"
	id, err := InsertData(sql, u.Username, u.Password, u.Email)
	u.Id = id
	return err
}
func (u *User) update() error {
	sql := "Update users set username=?, password=?, email=?"
	_, err := Exec(sql, u.Username, u.Password, u.Email)
	return err
}

func (u *User) Delete() error {
	sql := "Delete from users where id=?"
	_, err := Exec(sql, u.Id)
	return err
}

func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "select id,username,password,email, created_date from users where id=?"
	rows, _ := Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
	}

	return user
}

func GetUserByUsername(username string) *User {
	user := NewUser("", "", "")
	sql := "select id,username,password,email, created_date from users where username=?"
	rows, _ := Query(sql, username)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
	}

	return user
}

func Login(username, password string) bool {
	user := GetUserByUsername(username)
	fmt.Println(user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil

}

func GetUsers() *[]User {
	sql := "select id,username,password,email, created_date from users"
	var users []User
	rows, _ := Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
		users = append(users, user)
	}

	return &users
}

func ValidEmail(email string) bool {
	return emailRegexp.MatchString(email)
}
