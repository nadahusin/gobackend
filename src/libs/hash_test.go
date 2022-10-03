package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashingPassword(t *testing.T) {
	Pass := "password"
	result, err := HashingPassword(Pass)

	if err != nil {
		t.Fatal(err)
	}

	if result == Pass {
		t.Fatal(err, "failed to hash password")
	}

	assert.IsType(t, "string", result, "password is not a string")
}

func TestCheckPassword(t *testing.T) {
	err := CheckPassword("abcd123", "pass")

	assert.NotNil(t, err, "success to hash password")
}
