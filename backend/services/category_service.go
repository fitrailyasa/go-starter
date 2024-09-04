package services

import (
	"errors"
	"go-starter/backend/models"
)

var categories = []models.Category{}
var categoryID = 1

func GetAllCategories() []models.Category {
	return categories
}

func GetCategoryByID(id int) (*models.Category, error) {
	for _, category := range categories {
		if category.ID == id {
			return &category, nil
		}
	}
	return nil, errors.New("category not found")
}

func CreateCategory(name, description, img string) models.Category {
	category := models.Category{
		ID:          categoryID,
		Name:        name,
		Description: description,
		Images:      img,
	}
	categories = append(categories, category)
	categoryID++
	return category
}

func UpdateCategory(id int, name, description, img string) (*models.Category, error) {
	for i, category := range categories {
		if category.ID == id {
			categories[i].Name = name
			categories[i].Description = description
			categories[i].Images = img
			return &categories[i], nil
		}
	}
	return nil, errors.New("category not found")
}

func DeleteCategory(id int) error {
	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}
