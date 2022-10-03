package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	result := HelloName("jesen")

	if result != "HELLO jesen" {
		t.Fatal("return must be HELLO jesen")
	}
}

func TestHelloNames(t *testing.T) {
	result := HelloName("jesen")
	assert.Equal(t, "HELLO jesen", result, "return must be HELLO jesen")
}

func TestSubHelloName(t *testing.T) {
	t.Run("params nada", func(t *testing.T) {
		result := HelloName("nada")
		assert.Equal(t, "HELLO nada", result, "return must be HELLO nada")
	})
}
