package controllers

import (
	"go-starter/backend/models"
	"go-starter/backend/services"
	"html/template"
	"net/http"
	"strconv"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/user/index.html",
		"frontend/templates/pages/admin/user/create.html",
		"frontend/templates/pages/admin/user/edit.html",
		"frontend/templates/pages/admin/user/delete.html",
	))

	data := struct {
		Users []models.User
	}{
		Users: users,
	}

	tmpl.ExecuteTemplate(w, "app.html", data)
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "user" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := r.FormValue("role")
		status := r.FormValue("status")
		services.CreateUser(name, email, password, role, status)
		http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/user/index.html",
		"frontend/templates/pages/admin/user/create.html",
		"frontend/templates/pages/admin/user/edit.html",
		"frontend/templates/pages/admin/user/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	_, err := services.GetUserByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == "user" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := r.FormValue("role")
		status := r.FormValue("status")
		_, err = services.UpdateUser(id, name, email, password, role, status)
		if err == nil {
			http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/user/index.html",
		"frontend/templates/pages/admin/user/create.html",
		"frontend/templates/pages/admin/user/edit.html",
		"frontend/templates/pages/admin/user/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	services.DeleteUser(id)
	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}
