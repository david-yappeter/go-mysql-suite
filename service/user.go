package service

import (
	"errors"
	"myapp/config"
	"myapp/entity"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserCreate(input entity.User) (*entity.User, error) {
	input.Email = strings.ToLower(input.Email)
	_, err := UserGetByEmail(input.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == nil {
		return nil, errors.New("email not available")
	}

	db := config.GetDB()
	input.ID = uuid.New().String()
	if err := db.Model(input).Create(&input).Error; err != nil {
		return nil, err
	}

	return &input, nil
}

func UserGetAll() ([]*entity.User, error) {
	db := config.GetDB()

	var users []*entity.User
	if err := db.Model(users).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UserGetByEmail(email string) (*entity.User, error) {
	db := config.GetDB()

	var user entity.User
	if err := db.Model(user).Where("email like ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
