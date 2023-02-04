package infrastructures

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/infrastructures/rest"
	"github.com/naufalfmm/aquafarm-management-service/middlewares"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Infrastructures struct {
	Rest rest.Rest
}

func Init(usec usecases.Usecases, res resources.Resources, middl middlewares.Middlewares) (Infrastructures, error) {
	re, err := rest.Init(usec, res, middl)
	if err != nil {
		return Infrastructures{}, err
	}

	return Infrastructures{
		Rest: re,
	}, nil
}

func (i *Infrastructures) Register(ec *echo.Echo) {
	i.Rest.Register(ec)
}
