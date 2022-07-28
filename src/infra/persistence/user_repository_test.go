package persistence

import (
	"testing"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	conn, err := DBConf()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	// testCase := []struct{}

	var user = entity.User{}
	user.Email = "steven@example.com"
	user.FirstName = "victoria"
	user.LastName = "steven"
	user.Password = "password"

	repo := NewUserRepository(conn)

	t.Run("CreateUser", func(t *testing.T) {
		u, saveErr := repo.SaveUser(user)
		assert.Nil(t, saveErr)
		assert.EqualValues(t, u.Email, "steven@example.com")
		assert.EqualValues(t, u.FirstName, "victoria")
		assert.EqualValues(t, u.LastName, "steven")
		//The pasword is supposed to be hashed, so, it should not the same the one we passed:
		assert.NotEqual(t, u.Password, "password")
		assert.Equal(t, u.ID, uint64(1))

	})

	t.Run("GetUser", func(t *testing.T) {
		u, getUserErr := repo.GetUser(uint64(1))
		assert.Nil(t, getUserErr)
		assert.EqualValues(t, u.Email, "steven@example.com")
		assert.EqualValues(t, u.FirstName, "victoria")
		assert.EqualValues(t, u.LastName, "steven")

		assert.Equal(t, u.ID, uint64(1))

	})

}
