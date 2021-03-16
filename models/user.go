package models

import (
	"errors"
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
    password varchar(60) not null,
    email varchar(40),
    created_date timestamp default current_timestamp)`

func NewUser(username, password, email string) (*User, error) {
	user := &User{Username: username, Password: password, Email: email}
	if err := user.Valid(); err != nil {
		return &User{}, err
	}
	err := user.SetPassword(password)
	return user, err
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("No es posible cifrar el password")
	}
	u.Password = string(hash)
	return nil
}

func CreateUser(username, password, email string) (*User, error) {
	user, err := NewUser(username, password, email)
	if err != nil {
		return &User{}, err
	}
	err = user.Save()
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
	sql := "Update users set username=?, password=?, email=? where id=?"
	_, err := Exec(sql, u.Username, u.Password, u.Email, u.Id)
	return err
}

func (u *User) Delete() error {
	sql := "Delete from users where id=?"
	_, err := Exec(sql, u.Id)
	return err
}

func GetUser(sql string, conditional interface{}) *User {
	user := &User{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return user
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
	}

	return user
}

func GetUserById(id int) *User {
	sql := "select id,username,password,email, created_date from users where id=?"
	return GetUser(sql, id)
}

func GetUserByUsername(username string) *User {
	sql := "select id,username,password,email, created_date from users where username=?"
	return GetUser(sql, username)
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

func ValidEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("Formato de email invalido")
	}
	return nil
}
func (u *User) Valid() error {
	if err := ValidEmail(u.Email); err != nil {
		return err
	}

	if err := ValidUsername(u.Username); err != nil {
		return err
	}

	return nil
}

func ValidUsername(username string) error {
	if len(username) > 30 {
		return errors.New("username muy largo, maximo 30 caracteres")
	}
	return nil
}
