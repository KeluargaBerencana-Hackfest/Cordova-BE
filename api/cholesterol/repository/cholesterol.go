package repository

import (
	"context"

	"github.com/Ndraaa15/cordova/config/database"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/jmoiron/sqlx"
)

type CholesterolRepositoryImpl interface {
	SavedRecordCholesterol(c, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error)
	GetCholesterolHistory(c context.Context, id string) ([]*domain.CholesterolDB, error)
}

type CholesterolRepository struct {
	db *database.ClientDB
}

func NewCholesterolRepository(db *database.ClientDB) CholesterolRepositoryImpl {
	return &CholesterolRepository{db}
}

func (cr *CholesterolRepository) GetCholesterolHistory(c context.Context, id string) ([]*domain.CholesterolDB, error) {
	var cholesterol []*domain.CholesterolDB

	argKV := map[string]interface{}{
		"user_id": id,
	}

	query, args, err := sqlx.Named(GetCholesterolHistory, argKV)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = cr.db.Rebind(query)

	rows, err := cr.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cholesterolDB domain.CholesterolDB
		err := rows.StructScan(&cholesterolDB)
		if err != nil {
			return nil, err
		}

		cholesterol = append(cholesterol, &cholesterolDB)
	}

	return cholesterol, nil
}

func (cr *CholesterolRepository) SavedRecordCholesterol(c, id string, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error) {
	// get user by id
	// get percentage by hit model
	// by cholesterol level get recommendation
	return nil, nil
}
