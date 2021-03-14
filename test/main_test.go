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
	createDefaultUser()
}

func createDefaultUser() {
	sql := fmt.Sprintf("insert users set id='%d', username = '%s', password= '%s', email ='%s'", id, username, password, email)
	_, err := models.Exec(sql)
	if err != nil {
		panic(err)
	}

}

func afterTest() {
	fmt.Println(">> Despues de las pruebas")
	models.CloseConnection()
}
