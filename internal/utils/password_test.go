package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPasswordsEquality(t *testing.T) {
	pwd1 := "password123"
	pwd2 := "password123"
	assert.True(t, CheckPasswordsEquality(pwd1, pwd2))

	pwd1 = "password123"
	pwd2 = "anotherPassword"
	assert.False(t, CheckPasswordsEquality(pwd1, pwd2))
}

func TestCheckPasswordSecurity(t *testing.T) {
	// Checking function for checking password requirements
	//
	// Requirements: Min 8 symbols, min 1 UpperCase char, min 1 LowerCase char, min 1 Digit

	pwd := "Sh0rt"
	err := CheckPasswordSecurity(pwd)
	assert.Error(t, err)

	pwd = ""
	err = CheckPasswordSecurity(pwd)
	assert.Error(t, err)

	pwd = "s0mepassword"
	err = CheckPasswordSecurity(pwd)
	assert.Error(t, err)

	pwd = "123456789"
	err = CheckPasswordSecurity(pwd)
	assert.Error(t, err)

	pwd = "S0MEPASSWORD"
	err = CheckPasswordSecurity(pwd)
	assert.Error(t, err)

	pwd = "SomePassword"
	err = CheckPasswordSecurity(pwd)
	assert.Error(t, err)

	pwd = "Exactly8"
	err = CheckPasswordSecurity(pwd)
	assert.NoError(t, err)

	pwd = "L0ngpassword"
	err = CheckPasswordSecurity(pwd)
	assert.NoError(t, err)
}
