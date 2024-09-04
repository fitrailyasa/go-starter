package controllers

import (
	"go-starter/backend/services"
	"html/template"
	"net/http"
	"strconv"
)

func ProductController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/product/index.html",
		"frontend/templates/pages/admin/product/create.html",
		"frontend/templates/pages/admin/product/edit.html",
		"frontend/templates/pages/admin/product/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func CreateProductController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		img := r.FormValue("img")
		priceStr := r.FormValue("price")
		stockStr := r.FormValue("stock")
		categoryIDStr := r.FormValue("category_id")

		price, err := strconv.Atoi(priceStr)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
			return
		}
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		services.CreateProduct(name, description, img, price, stock, categoryID)
		http.Redirect(w, r, "/admin/product", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/product/index.html",
		"frontend/templates/pages/admin/product/create.html",
		"frontend/templates/pages/admin/product/edit.html",
		"frontend/templates/pages/admin/product/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func UpdateProductController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	_, err := services.GetProductByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		img := r.FormValue("img")
		priceStr := r.FormValue("price")
		stockStr := r.FormValue("stock")
		categoryIDStr := r.FormValue("category_id")

		price, err := strconv.Atoi(priceStr)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
			return
		}
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		_, err = services.UpdateProduct(id, name, description, img, price, stock, categoryID)
		if err == nil {
			http.Redirect(w, r, "/admin/product", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles(
		"frontend/templates/layouts/admin/app.html",
		"frontend/templates/layouts/admin/navbar.html",
		"frontend/templates/layouts/admin/sidebar.html",
		"frontend/templates/layouts/admin/footer.html",
		"frontend/templates/pages/admin/product/index.html",
		"frontend/templates/pages/admin/product/create.html",
		"frontend/templates/pages/admin/product/edit.html",
		"frontend/templates/pages/admin/product/delete.html",
	))
	tmpl.ExecuteTemplate(w, "app.html", nil)
}

func DeleteProductController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	services.DeleteProduct(id)
	http.Redirect(w, r, "/admin/product", http.StatusSeeOther)
}
