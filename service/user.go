package service

import (
	"context"
	"goapi/config"
	"goapi/graph/model"
	"goapi/tools"
	"strings"
)

func UserCreate(ctx context.Context, input model.InputUser) (*model.User, error) {
	db := config.GetDB()

	input.Password = tools.HashPassword(input.Password)

	user := model.User{
		FirstName: input.Firstname,
		LastName:  input.Lastname,
		Email:     strings.ToLower(input.Email),
		Password:  input.Password,
	}

	if err := db.Model(user).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetByID(ctx context.Context, id int64) (*model.User, error) {
	db := config.GetDB()

	var user model.User
	if err := db.Model(user).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := config.GetDB()

	var user model.User
	if err := db.Model(user).Where("email LIKE ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
