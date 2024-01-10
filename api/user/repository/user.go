package repository

import "github.com/Ndraaa15/cordova/config/database"

type UserRepositoryImpl interface {
}

type UserRepository struct {
	db *database.ClientDB
}

func NewUserRepository(db *database.ClientDB) UserRepositoryImpl {
	return &UserRepository{db}
}
