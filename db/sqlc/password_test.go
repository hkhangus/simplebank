package db

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(6)

	hashedPass1, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass1)

	err = util.CheckPassword(hashedPass1, password)
	require.NoError(t, err)

	wrongPassword := util.RandomString(6)
	err = util.CheckPassword(hashedPass1, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPass2, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass2)
	require.NotEqual(t, hashedPass1, hashedPass2)
}