package controllers

import (
	farmsControllers "github.com/naufalfmm/aquafarm-management-service/infrastructures/rest/controllers/farms"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Controllers struct {
	Farms farmsControllers.Controllers
}

func Init(usec usecases.Usecases, res resources.Resources) (Controllers, error) {
	farmsCont, err := farmsControllers.Init(usec, res)
	if err != nil {
		return Controllers{}, err
	}

	return Controllers{
		Farms: farmsCont,
	}, nil
}
