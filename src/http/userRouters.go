package http

import (
	"github.com/andrersp/financeiro-go/src/application/api"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"github.com/andrersp/financeiro-go/src/infra/auth"
)

func loadUserRouters(userService repository.UserRepository) []Routers {
	token := auth.NewToken()
	user := api.NewUserApi(userService, token)
	routers := []Routers{
		{
			Method: "POST",
			URI:    "/users",
			Func:   user.SaveUser,
		},
		{
			Method: "GET",
			URI:    "/users",
			Func:   user.GetUsers,
		},
		{
			Method: "GET",
			URI:    "/users/:userID",
			Func:   user.GetUser,
		},
	}
	return routers
}
