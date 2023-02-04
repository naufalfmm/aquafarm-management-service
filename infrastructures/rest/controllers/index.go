package controllers

import (
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type Controllers struct{}

func Init(usec usecases.Usecases, res resources.Resources) (Controllers, error) {
	return Controllers{}, nil
}
