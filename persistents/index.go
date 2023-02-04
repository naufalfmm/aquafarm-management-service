package persistents

import (
	"github.com/naufalfmm/aquafarm-management-service/persistents/repositories"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type Persistents struct {
	Repositories repositories.Repositories
}

func Init(res resources.Resources) (Persistents, error) {
	reps, err := repositories.Init(res)
	if err != nil {
		return Persistents{}, err
	}

	return Persistents{
		Repositories: reps,
	}, nil
}
