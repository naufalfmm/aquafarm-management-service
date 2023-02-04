package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/infrastructures/rest/routes"
	"github.com/naufalfmm/aquafarm-management-service/middlewares"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Rest struct {
	Routes routes.Routes
}

func Init(usec usecases.Usecases, res resources.Resources, middl middlewares.Middlewares) (Rest, error) {
	rout, err := routes.Init(usec, res, middl)
	if err != nil {
		return Rest{}, err
	}

	return Rest{rout}, nil
}

func (r *Rest) Register(ec *echo.Echo) {
	r.Routes.Register(ec)
}
