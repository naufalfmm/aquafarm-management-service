package hooks

import (
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Hooks struct {
	Usecases  usecases.Usecases
	Resources resources.Resources
}

func Init(usecases usecases.Usecases, resources resources.Resources) (Hooks, error) {
	return Hooks{
		Usecases:  usecases,
		Resources: resources,
	}, nil
}
