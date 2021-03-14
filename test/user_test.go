package test

import (
	"fmt"
	"math/rand"
	"rest/models"
	"testing"
	"time"
)

var user *models.User

const (
	id          = 1
	username    = "diego"
	password    = "123456"
	email       = "di_564@hotmail.com"
	createdDate = "2017-08-17"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser(username, password, email)

	if user.Username != username || user.Password != password || user.Email != email {
		t.Error("No es posible crear el objeto")
	}
}

func TestSave(t *testing.T) {
	user := models.NewUser(randomUserName(), password, email)
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(randomUserName(), password, email)

	if err != nil {
		t.Error("No es posible insertar el objeto", nil)
	}
}

func randomUserName() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}

func TestUniqueUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)

	if err == nil {
		t.Error("Es posible insertar registros con usernames duplicados")
	}
}

func TestDuplicateUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'users.username'", username)
	if err.Error() != message {
		t.Error("Username duplicado")
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUser(id)
	if equalsUser(user) {
		t.Error("No es posible obtener el usuario")
	}
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(*users) == 0 {
		t.Error("No es posible obtener a los usuarios")
	}
}

func TestDeleteUser(t *testing.T) {
	if err := user.Delete(); err != nil {
		t.Error("No es posible eliminar al usuario")
	}
}

func equalsUser(user *models.User) bool {
	return user.Username != username && user.Password != password && user.Email != email
}

func equalsCreatedDate(date time.Time) bool {
	t, _ := time.Parse("20016-01-02", createdDate)
	return t == date

}
