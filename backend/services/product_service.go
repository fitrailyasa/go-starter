package services

import (
	"errors"
	"go-starter/backend/db"
	"go-starter/backend/models"
)

var products = []models.Product{}
var productID = 1

func GetAllProducts() ([]models.Product, error) {
	rows, err := db.DB.Query("SELECT id, name, description, img, price, stock, category_id FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Images, &product.Price, &product.Stock, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
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
		ID:          productID,
		Name:        name,
		Description: description,
		Images:      img,
		Price:       price,
		Stock:       stock,
		CategoryID:  category_id,
	}
	products = append(products, product)
	productID++
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
