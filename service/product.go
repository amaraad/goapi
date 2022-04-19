package service

import (
	"context"
	"goapi/config"
	"goapi/graph/model"
)

func AddProduct(ctx context.Context, product model.InputProduct) (interface{}, error) {
	db := config.GetDB()
	prod := model.Product{}
	prod.Name = product.Name
	prod.Category.ID = product.CategoryID
	if err := db.Create(&prod).Error; err != nil {
		return nil, err
	} else {
		return success, nil
	}
}

func EditProduct(ctx context.Context, product model.InputProduct, productID int64) (interface{}, error) {
	db := config.GetDB()
	prod := model.Product{}
	db.First(&prod, productID)
	prod.Name = product.Name
	prod.Category.ID = product.CategoryID
	if err := db.Save(&prod).Error; err != nil {
		return nil, err
	} else {
		return success, nil
	}
}

func DeleteProduct(id int64) (interface{}, error) {
	db := config.GetDB()
	prod := model.Product{}
	db.First(&prod, id)
	if err := db.Delete(&prod).Error; err != nil {
		return nil, err
	} else {
		return success, nil
	}
}

func Products(ctx context.Context, search *string, categoryID *int64) ([]*model.Product, error) {
	db := config.GetDB()
	var products []*model.Product
	var query = "name like ?"
	if categoryID != nil && search != nil {
		query = query + " and category_id = ?"
		db.Model(&model.Product{}).Where(query, "%"+*search+"%", categoryID).Find(&products)
	} else if categoryID != nil {
		query = "category_id = ?"
		db.Model(&model.Product{}).Where(query, categoryID).Find(&products)
	} else if search != nil {
		db.Model(&model.Product{}).Where(query, "%"+*search+"%").Find(&products)
	} else {
		db.Model(&model.Product{}).Find(&products)
	}

	return products, nil
}
