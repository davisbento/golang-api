package repository_test

import (
	"davisbento/golang-api/db"
	"davisbento/golang-api/db/entity"
	"davisbento/golang-api/db/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAll(t *testing.T) {
	conn, err := db.NewDB().Connect("../../../sqlite/test.db")
	require.NoError(t, err)

	userRepo := repository.NewUserRepository(conn)
	var users []*entity.User

	users, err = userRepo.FindAll()

	require.NoError(t, err)
	require.NotEmpty(t, users)
}
