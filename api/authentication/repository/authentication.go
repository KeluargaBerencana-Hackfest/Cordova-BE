package repository

import "github.com/Ndraaa15/cordova/config/database"

type AuthRepositoryImpl interface {
}

type AuthRepository struct {
	db *database.ClientDB
}

func NewAuthRepository(db *database.ClientDB) AuthRepositoryImpl {
	return &AuthRepository{db}
}
