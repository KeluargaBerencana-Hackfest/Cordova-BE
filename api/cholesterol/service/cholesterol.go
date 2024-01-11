package service

import (
	"context"

	"github.com/Ndraaa15/cordova/api/cholesterol/repository"
	"github.com/Ndraaa15/cordova/domain"
)

type CholesterolServiceImpl interface {
	CalculateCholesterol(ctx context.Context)
	GetCholesterolHistory(ctx context.Context, id string) (*domain.Cholesterol, error)
}

type CholesterolService struct {
	cr repository.CholesterolRepositoryImpl
}

func NewCholesterolService(cholesterolRepository repository.CholesterolRepositoryImpl) CholesterolServiceImpl {
	return &CholesterolService{cholesterolRepository}
}

func (cs *CholesterolService) CalculateCholesterol(ctx context.Context) {

}

func (cs *CholesterolService) GetCholesterolHistory(ctx context.Context, id string) (*domain.Cholesterol, error) {
	cholesterols := &domain.Cholesterol{}
	cholesterolMap := make(map[uint64][]*domain.CholesterolDB)

	cholesterol, err := cs.cr.GetCholesterolHistory(id)
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
