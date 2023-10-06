package repository

import (
	"fmt"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/model"
	"gorm.io/gorm"
)

type User interface {
	Save(user model.User) (model.User, error)
	FindByID(id uint) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type user struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) User {
	return &user{
		db: db,
	}
}

func (u user) Save(user model.User) (model.User, error) {
	u.db.Create(&user)
	return user, nil
}

func (u user) FindByID(id uint) (model.User, error) {
	var user model.User
	u.db.First(&user, id)
	if user.ID == 0 {
		return model.User{}, dto.AppError(fmt.Errorf("user with id %d not found", id))
	}
	return user, nil
}

func (u user) FindByEmail(email string) (model.User, error) {
	var user model.User
	u.db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return model.User{}, dto.AppError(fmt.Errorf("user with email %s not found", email))
	}
	return user, nil
}
