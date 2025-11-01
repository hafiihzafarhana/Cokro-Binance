package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Dummy struct untuk test
type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=18,lte=99"`
}

func TestValidateStruct_Success(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	err := ValidateStruct(user)
	assert.NoError(t, err)
}

func TestValidateStruct_ErrorRequired(t *testing.T) {
	user := User{
		Name:  "", // missing required
		Email: "john@example.com",
		Age:   30,
	}

	err := ValidateStruct(user)
	assert.Error(t, err) // error
}

func TestValidateStruct_ErrorEmail(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "not-an-email", // invalid email
		Age:   30,
	}

	err := ValidateStruct(user)
	assert.Error(t, err) // error
}

func TestValidateStruct_ErrorAge(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   10, // age < 18
	}

	err := ValidateStruct(user)
	assert.Error(t, err) // error
}