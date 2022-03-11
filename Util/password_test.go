package Util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	passowrd := RandomStringGenerator(6)

	hashedPassword, err := HashPassword(passowrd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(passowrd, hashedPassword)
	require.NoError(t, err)

	wrongPassowrd := RandomStringGenerator(6)
	err = CheckPassword(wrongPassowrd, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword1, err := HashPassword(passowrd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)
	require.NotEqual(t, hashedPassword, hashedPassword1)

	err = CheckPassword(passowrd, hashedPassword1)
	require.NoError(t, err)
}
