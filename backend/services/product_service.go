package services

import (
	"errors"
	"go-starter/backend/models"
)

var products = []models.Product{}
var product_id = 1

func GetAllProducts() []models.Product {
	return products
}

func GetProductByID(id int) (*models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func CreateProduct(name, description, img string, price, stock, category_id int) models.Product {
	product := models.Product{
		ID:          product_id,
		Name:        name,
		Description: description,
		Images:      img,
		Price:       price,
		Stock:       stock,
		CategoryID:  category_id,
	}
	products = append(products, product)
	product_id++
	return product
}

func UpdateProduct(id int, name, description, img string, price, stock, category_id int) (*models.Product, error) {
	for i, product := range products {
		if product.ID == id {
			products[i].Name = name
			products[i].Description = description
			products[i].Images = img
			products[i].Price = price
			products[i].Stock = stock
			products[i].CategoryID = category_id
			return &products[i], nil
		}
	}
	return nil, errors.New("product not found")
}

func DeleteProduct(id int) error {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
