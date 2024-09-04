package db

import (
	"log"
)

func SeedData() {
	// Dummy data for user table
	users := []struct {
		name     string
		email    string
		password string
		role     string
		status   string
	}{
		{"Alice", "alice@example.com", "password123", "admin", "active"},
		{"Bob", "bob@example.com", "password123", "user", "inactive"},
		{"Charlie", "charlie@example.com", "password123", "user", "active"},
	}

	// Insert dummy users
	for _, user := range users {
		_, err := DB.Exec(`INSERT INTO user (name, email, password, role, status) VALUES (?, ?, ?, ?, ?)`,
			user.name, user.email, user.password, user.role, user.status)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Dummy data for category table
	categories := []struct {
		name        string
		description string
		img         string
	}{
		{"Electronics", "Devices and gadgets", "electronics.jpg"},
		{"Clothing", "Apparel and accessories", "clothing.jpg"},
		{"Home & Kitchen", "Furniture and kitchenware", "home_kitchen.jpg"},
	}

	// Insert dummy categories
	for _, category := range categories {
		_, err := DB.Exec(`INSERT INTO category (name, description, img) VALUES (?, ?, ?)`,
			category.name, category.description, category.img)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Dummy data for product table
	products := []struct {
		name        string
		description string
		img         string
		price       int
		stock       int
		categoryID  int
	}{
		{"Laptop", "High-performance laptop", "laptop.jpg", 1000, 10, 1},
		{"T-shirt", "Cotton t-shirt", "tshirt.jpg", 20, 100, 2},
		{"Blender", "High-speed blender", "blender.jpg", 50, 30, 3},
	}

	// Insert dummy products
	for _, product := range products {
		_, err := DB.Exec(`INSERT INTO product (name, description, img, price, stock, category_id) VALUES (?, ?, ?, ?, ?, ?)`,
			product.name, product.description, product.img, product.price, product.stock, product.categoryID)
		if err != nil {
			log.Fatal(err)
		}
	}
}
