package interfaces

import (
	"github.com/andrersp/financeiro-go/src/infra/persistence"
	"github.com/andrersp/financeiro-go/src/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	URI         string
	Func        func(c *gin.Context)
	Method      string
	RequireAuth bool
}

func loadRouters(services persistence.Repositories, r *gin.Engine) *gin.Engine {
	routers := []Routers{}
	userRouters := loadUserRouters(services.User)

	routers = append(routers, userRouters...)

	v1 := r.Group("v1")

	for _, router := range routers {

		if router.RequireAuth {
			v1.Handle(router.Method, router.URI, middleware.AuthMiddleware(), router.Func)

		} else {
			v1.Handle(router.Method, router.URI, router.Func)

		}

	}

	return r
}

func StartServer(services persistence.Repositories) *gin.Engine {
	r := gin.Default()

	return loadRouters(services, r)

}
