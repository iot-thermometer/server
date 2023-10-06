package repository

import "gorm.io/gorm"

type Reading interface {
}

type reading struct {
	db *gorm.DB
}

func newReadingRepository(db *gorm.DB) Reading {
	return &reading{
		db: db,
	}
}
