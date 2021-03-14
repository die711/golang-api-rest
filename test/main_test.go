package test

import (
	"fmt"
	"os"
	"rest/models"
	"testing"
)

func TestMain(m *testing.M) {
	beforeTest()
	result := m.Run()
	afterTest()
	os.Exit(result)

}

func beforeTest() {
	fmt.Println(">> Antes de las pruebas")
	models.CreateConnection()
}

func afterTest() {
	fmt.Println(">> Despues de las pruebas")
	models.CloseConnection()
}
