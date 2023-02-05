package controllers

import (
	farmsControllers "github.com/naufalfmm/aquafarm-management-service/infrastructures/rest/controllers/farms"
	pondsControllers "github.com/naufalfmm/aquafarm-management-service/infrastructures/rest/controllers/ponds"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Controllers struct {
	Farms farmsControllers.Controllers
	Ponds pondsControllers.Controllers
}

func Init(usec usecases.Usecases, res resources.Resources) (Controllers, error) {
	farmsCont, err := farmsControllers.Init(usec, res)
	if err != nil {
		return Controllers{}, err
	}

	pondsCont, err := pondsControllers.Init(usec, res)
	if err != nil {
		return Controllers{}, err
	}

	return Controllers{
		Farms: farmsCont,
		Ponds: pondsCont,
	}, nil
}
