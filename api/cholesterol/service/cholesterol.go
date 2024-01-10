package service

import (
	"context"

	"github.com/Ndraaa15/cordova/api/cholesterol/repository"
)

type CholesterolServiceImpl interface {
	CalculateCholesterol(ctx context.Context)
}

type CholesterolService struct {
	cr repository.CholesterolRepositoryImpl
}

func NewCholesterolService(cholesterolRepository *repository.CholesterolRepositoryImpl) CholesterolServiceImpl {
	return &CholesterolService{cholesterolRepository}
}

func (cs *CholesterolService) CalculateCholesterol(ctx context.Context) {

}
