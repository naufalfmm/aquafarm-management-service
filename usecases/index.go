package usecases

import (
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type Usecases struct{}

func Init(persist persistents.Persistents, res resources.Resources) (Usecases, error) {
	return Usecases{}, nil
}
