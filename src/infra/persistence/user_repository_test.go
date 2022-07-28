package persistence

import (
	"fmt"
	"strings"
	"testing"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	conn, err := DBConf()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	testCases := []struct {
		User entity.User
	}{
		{
			User: entity.User{
				FirstName: "victoria",
				LastName:  "steven",
				Email:     "steven@example.com",
				Password:  "password",
			},
		},
	}

	repo := NewUserRepository(conn)

	t.Run("CreateUserSuccess", func(t *testing.T) {

		for index, testCase := range testCases {
			u, saveErr := repo.SaveUser(testCase.User)
			assert.Nil(t, saveErr)
			assert.EqualValues(t, u.Email, testCase.User.Email)
			assert.EqualValues(t, u.FirstName, testCase.User.FirstName)
			assert.EqualValues(t, u.LastName, testCase.User.LastName)
			//The pasword is supposed to be hashed, so, it should not the same the one we passed:
			assert.NotEqual(t, u.Password, testCase.User.Password)
			assert.Equal(t, u.ID, uint64(index+1))

		}

	})

	t.Run("CreateUserErrorDuplicateEEmail", func(t *testing.T) {

		for _, testCase := range testCases {
			u, saveErr := repo.SaveUser(testCase.User)
			assert.Nil(t, u)

			fmt.Printf(saveErr.Error())

			assert.Equal(t, true, strings.Contains(saveErr.Error(), "UNIQUE"))

		}

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
