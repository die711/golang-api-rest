package test

import "testing"

func TestHolaMundo(t *testing.T) {
	str := "Hola Mundo"
	if str != "Hola Mundo" {
		t.Error("No es posbile saludar a los usuariops", nil)
	}
}
