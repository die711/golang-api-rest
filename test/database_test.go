package test

import (
	"rest/models"
	"testing"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()

	if connection == nil {
		t.Error("No es posible realizar la conexion", nil)
	}
}
