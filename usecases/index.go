package usecases

import (
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	farmsUsecases "github.com/naufalfmm/aquafarm-management-service/usecases/farms"
)

type Usecases struct {
	Farms farmsUsecases.Usecases
}

func Init(persist persistents.Persistents, res resources.Resources) (Usecases, error) {
	farmUsec, err := farmsUsecases.Init(persist, res)
	if err != nil {
		return Usecases{}, err
	}

	return Usecases{
		Farms: farmUsec,
	}, nil
}
