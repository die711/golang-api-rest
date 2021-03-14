package test

import (
	"fmt"
	"os"
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
}

func afterTest() {
	fmt.Println(">> Despues de las pruebas")
}
