package usecases

import (
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	farmsUsecases "github.com/naufalfmm/aquafarm-management-service/usecases/farms"
	pondsUsecases "github.com/naufalfmm/aquafarm-management-service/usecases/ponds"
)

type Usecases struct {
	Farms farmsUsecases.Usecases
	Ponds pondsUsecases.Usecases
}

func Init(persist persistents.Persistents, res resources.Resources) (Usecases, error) {
	farmUsec, err := farmsUsecases.Init(persist, res)
	if err != nil {
		return Usecases{}, err
	}

	pondUsec, err := pondsUsecases.Init(persist, res)
	if err != nil {
		return Usecases{}, err
	}

	return Usecases{
		Farms: farmUsec,
		Ponds: pondUsec,
	}, nil
}
