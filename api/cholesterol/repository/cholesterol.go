package repository

import "github.com/Ndraaa15/cordova/config/database"

type CholesterolRepositoryImpl interface {
}

type CholesterolRepository struct {
	db *database.ClientDB
}

func NewCholesterolRepository(db *database.ClientDB) CholesterolRepositoryImpl {
	return &CholesterolRepository{db}
}
