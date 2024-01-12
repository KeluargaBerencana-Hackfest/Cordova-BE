package service

import (
	"context"
	"net/http"

	"github.com/Ndraaa15/cordova/api/cholesterol/repository"
	"github.com/Ndraaa15/cordova/domain"
)

type CholesterolServiceImpl interface {
	CalculateCholesterol(c context.Context, req *domain.CholesterolRequest) (*domain.Cholesterol, error)
	GetCholesterolHistory(c context.Context, id string) (*domain.Cholesterol, error)
}

type CholesterolService struct {
	cr repository.CholesterolRepositoryImpl
}

func NewCholesterolService(cholesterolRepository repository.CholesterolRepositoryImpl) CholesterolServiceImpl {
	return &CholesterolService{cholesterolRepository}
}

func (cs *CholesterolService) CalculateCholesterol(c context.Context, req *domain.CholesterolRequest) (*domain.Cholesterol, error) {
	//Deploying Using Google Cloud Platform
	http.Post("https://cordova-model-j5ofojnjyq-as.a.run.app", "application/json", nil)
	return nil, nil
}

func (cs *CholesterolService) GetCholesterolHistory(c context.Context, id string) (*domain.Cholesterol, error) {
	cholesterols := &domain.Cholesterol{}
	cholesterolMap := make(map[uint64][]*domain.CholesterolDB)

	cholesterol, err := cs.cr.GetCholesterolHistory(c, id)
	if err != nil {
		return nil, err
	}

	for _, value := range cholesterol {
		cholesterolMap[value.Year] = append(cholesterolMap[value.Year], value)
	}

	cholesterols.UserID = id
	cholesterols.Cholesterols = cholesterolMap

	return cholesterols, nil
}
