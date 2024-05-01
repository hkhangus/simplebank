package db

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(6)

	hashedPass, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass)

	err = util.CheckPassword(hashedPass, password)
	require.NoError(t, err)

	wrongPassword := util.RandomString(6)
	err = util.CheckPassword(hashedPass, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}