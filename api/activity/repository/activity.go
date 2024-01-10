package repository

import "github.com/Ndraaa15/cordova/config/database"

type ActivityRepositoryImpl interface {
}

type ActivityRepository struct {
	db *database.ClientDB
}

func NewActivitylRepository(db *database.ClientDB) ActivityRepositoryImpl {
	return &ActivityRepository{db}
}
