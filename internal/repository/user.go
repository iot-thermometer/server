package repository

import "gorm.io/gorm"

type User interface {
}

type user struct {
	db gorm.DB
}

func newUserRepository(db gorm.DB) User {
	return &user{
		db: db,
	}
}
