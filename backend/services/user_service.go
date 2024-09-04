package services

import (
	"errors"
	"go-starter/backend/db"
	"go-starter/backend/models"
)

var users = []models.User{}
var userID = 1

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, name, email, password, role, status FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func CreateUser(name, email, password, role, status string) models.User {
	user := models.User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
		Status:   status,
	}
	users = append(users, user)
	userID++
	return user
}

func UpdateUser(id int, name, email, password, role, status string) (*models.User, error) {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			users[i].Email = email
			users[i].Password = password
			users[i].Role = role
			users[i].Status = status
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func DeleteUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
