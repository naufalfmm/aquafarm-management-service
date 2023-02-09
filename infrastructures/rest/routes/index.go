package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/infrastructures/rest/controllers"
	"github.com/naufalfmm/aquafarm-management-service/middlewares"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"

	_ "github.com/naufalfmm/aquafarm-management-service/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	ec.Pre(r.Middlewares.RequestLogger(), r.Middlewares.RemoveTrailingSlash())
	v1 := ec.Group("/v1")

	farm := v1.Group("/farms", r.Middlewares.VerifyToken())
	farm.GET("", r.Controllers.Farms.GetAllPaginated)
	farm.POST("", r.Controllers.Farms.Create)
	farm.DELETE("/:id", r.Controllers.Farms.DeleteByID)
	farm.GET("/:id", r.Controllers.Farms.GetByID)
	farm.PUT("/:code", r.Controllers.Farms.Upsert)
	farm.POST("/:id/ponds", r.Controllers.Ponds.Create)
	farm.PUT("/:id/ponds/:code", r.Controllers.Ponds.Upsert)

	pond := v1.Group("/ponds", r.Middlewares.VerifyToken())
	pond.GET("", r.Controllers.Ponds.GetAllPaginated)
	pond.POST("", r.Controllers.Ponds.Create)
	pond.DELETE("/:id", r.Controllers.Ponds.DeleteByID)
	pond.GET("/:id", r.Controllers.Ponds.GetByID)
	pond.PUT("/:code", r.Controllers.Ponds.Upsert)

	endpoint := v1.Group("/endpoints", r.Middlewares.VerifyToken())
	endpoint.GET("/reports", r.Controllers.EndpointLogs.GetAllReports)

	v1.GET("/docs/*", echoSwagger.WrapHandler)
}
