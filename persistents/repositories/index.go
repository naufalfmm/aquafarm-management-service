package repositories

import "github.com/naufalfmm/aquafarm-management-service/resources"

type Repositories struct {
}

func Init(res resources.Resources) (Repositories, error) {
	return Repositories{}, nil
}
