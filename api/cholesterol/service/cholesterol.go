package service

import (
	"github.com/Ndraaa15/cordova/api/cholesterol/repository"
)

type CholesterolServiceImpl interface {
}

type CholesterolService struct {
	cr repository.CholesterolRepositoryImpl
}

func NewCholesterolService(cholesterolRepository *repository.CholesterolRepositoryImpl) CholesterolServiceImpl {
	return &CholesterolService{cholesterolRepository}
}
