package test

import (
	"fmt"
	"math/rand"
	"rest/models"
	"testing"
)

const (
	username = "diego"
	password = "123456"
	email    = "di_564@hotmail.com"
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
