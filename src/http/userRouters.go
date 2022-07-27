package http

import (
	"github.com/andrersp/financeiro-go/src/application/api"
	"github.com/andrersp/financeiro-go/src/domain/repository"
)

func loadUserRouters(userService repository.UserRepository) []Routers {
	user := api.NewUserApi(userService)
	routers := []Routers{
		{
			Method: "POST",
			URI:    "/user",
			Func:   user.SaveUser,
		},
		{
			Method: "GET",
			URI:    "/user",
			Func:   user.GetUsers,
		},
		{
			Method: "GET",
			URI:    "/user/:userID",
			Func:   user.GetUser,
		},
	}
	return routers
}
