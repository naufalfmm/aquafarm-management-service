package farmsControllers

import (
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Controllers struct {
	Usecases  usecases.Usecases
	Resources resources.Resources
}

func Init(usecases usecases.Usecases, resources resources.Resources) (Controllers, error) {
	return Controllers{
		Usecases:  usecases,
		Resources: resources,
	}, nil
}
