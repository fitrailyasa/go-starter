package controllers

import (
	"go-starter/backend/models"
	"go-starter/backend/services"
	"html/template"
	"net/http"
	"strconv"
)

func CategoryController(w http.ResponseWriter, r *http.Request) {
	categories, err := services.GetAllCategories()
	if err != nil {
		http.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/category/index.html",
		"frontend/templates/pages/admin/category/create.html",
		"frontend/templates/pages/admin/category/edit.html",
		"frontend/templates/pages/admin/category/delete.html",
	))

	data := struct {
		Categories []models.Category
	}{
		Categories: categories,
	}

	tmpl.ExecuteTemplate(w, "app.html", data)
}

func CreateCategoryController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "category" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		img := r.FormValue("img")
		services.CreateCategory(name, description, img)
		http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/category/index.html",
		"frontend/templates/pages/admin/category/create.html",
		"frontend/templates/pages/admin/category/edit.html",
		"frontend/templates/pages/admin/category/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func UpdateCategoryController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	_, err := services.GetCategoryByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == "category" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		img := r.FormValue("img")
		_, err = services.UpdateCategory(id, name, description, img)
		if err == nil {
			http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/category/index.html",
		"frontend/templates/pages/admin/category/create.html",
		"frontend/templates/pages/admin/category/edit.html",
		"frontend/templates/pages/admin/category/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func DeleteCategoryController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	services.DeleteCategory(id)
	http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
}
