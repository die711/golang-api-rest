package models

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:password`
	Email    string `json:email`
}

const UserSchema string = `create table users(
	id int(6) auto_increment primary key,
    username varchar(30) not null,
    password varchar(64) not null,
    email varchar(40),
    created_date timestamp default current_timestamp)`

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
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
	result, err := Exec(sql, u.Username, u.Password, u.Email)
	u.Id, _ = result.LastInsertId()
	return err
}
func (u *User) update() error {
	sql := "Update users set username=?, password=?, email=?"
	_, err := Exec(sql, u.Username, u.Password, u.Email)
	return err
}

func (u *User) Delete() {
	sql := "Delete from users where id=?"
	Exec(sql, u.Id)
}

func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "select id,username,password,email from users where id=?"
	rows, _ := Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	return user

}

func GetUsers() *[]User {
	sql := "select id,username,password,email from users"
	var users []User

	rows, _ := Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return &users

}
