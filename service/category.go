package service

import (
	"context"
	"goapi/config"
	"goapi/graph/model"
)

var success = "OK"

func AddCategory(ctx context.Context, category model.InputCategory) (interface{}, error) {
	db := config.GetDB()
	cat := model.Category{}
	cat.Name = category.Name
	if err := db.Create(&cat).Error; err != nil {
		return nil, err
	} else {
		return success, nil
	}
}

func EditCategory(ctx context.Context, category model.InputCategory, categoryID int64) (interface{}, error) {
	db := config.GetDB()
	cat := model.Category{}
	db.First(&cat, categoryID)
	cat.Name = category.Name
	if err := db.Save(&cat).Error; err != nil {
		return nil, err
	} else {
		return success, nil
	}
}

func DeleteCategory(ctx context.Context, id int64) (interface{}, error) {
	db := config.GetDB()
	cat := model.Category{}
	db.First(&cat, id)
	if err := db.Delete(&cat).Error; err != nil {
		return nil, err
	} else {
		return success, nil
	}
}

func Categories(ctx context.Context, search *string) ([]*model.Category, error) {
	db := config.GetDB()
	var categories []*model.Category
	if search != nil {
		db.Model(&model.Category{}).Where("name like ?", "%"+*search+"%").Find(&categories)
	} else {
		db.Model(&model.Category{}).Find(&categories)
	}

	return categories, nil
}
