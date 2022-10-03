package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespone(t *testing.T) {
	result := Respone(nil, 200, false)

	assert.Nil(t, result.IsError, result.IsError)
	assert.Equal(t, 200, "status must be code 200")

}
