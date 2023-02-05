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
	v1 := ec.Group("/v1")

	farm := v1.Group("/farms", r.Middlewares.VerifyToken())
	farm.POST("", r.Controllers.Farms.Create)
	farm.POST("/:id/ponds", r.Controllers.Ponds.Create)

	pond := v1.Group("/ponds", r.Middlewares.VerifyToken())
	pond.POST("", r.Controllers.Ponds.Create)
	pond.DELETE("/:id", r.Controllers.Ponds.DeleteByID)
}
