package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/infrastructures/rest/controllers"
	"github.com/naufalfmm/aquafarm-management-service/middlewares"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Routes struct {
	Controllers controllers.Controllers
	Middlewares middlewares.Middlewares
}

func Init(usec usecases.Usecases, res resources.Resources, middlewares middlewares.Middlewares) (Routes, error) {
	cont, err := controllers.Init(usec, res)
	if err != nil {
		return Routes{}, err
	}

	return Routes{
		Controllers: cont,
		Middlewares: middlewares,
	}, nil
}

func (r *Routes) Register(ec *echo.Echo) {
	ec.Pre(r.Middlewares.RemoveTrailingSlash())
}
