package routes

import (
	"go-starter/backend/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// CLIENT SIDE //
	router.HandleFunc("/", controllers.HomeController)
	router.HandleFunc("/about", controllers.AboutController)
	router.HandleFunc("/contact", controllers.ContactController)

	// ADMIN SIDE //
	// CRUD USER
	router.HandleFunc("/admin/user", controllers.UserController).Methods("GET")
	router.HandleFunc("/admin/user", controllers.CreateUserController).Methods("POST")
	router.HandleFunc("/admin/user/{id:[0-9]+}", controllers.UpdateUserController).Methods("PUT")
	router.HandleFunc("/admin/user/{id:[0-9]+}", controllers.DeleteUserController).Methods("POST")

	// CRUD PRODUCT
	router.HandleFunc("/admin/product", controllers.ProductController).Methods("GET")
	router.HandleFunc("/admin/product", controllers.CreateProductController).Methods("POST")
	router.HandleFunc("/admin/product/{id:[0-9]+}", controllers.UpdateProductController).Methods("PUT")
	router.HandleFunc("/admin/product/{id:[0-9]+}", controllers.DeleteProductController).Methods("POST")

	// CRUD CATEGORY
	router.HandleFunc("/admin/category", controllers.CategoryController).Methods("GET")
	router.HandleFunc("/admin/category", controllers.CreateCategoryController).Methods("POST")
	router.HandleFunc("/admin/category/{id:[0-9]+}", controllers.UpdateCategoryController).Methods("PUT")
	router.HandleFunc("/admin/category/{id:[0-9]+}", controllers.DeleteCategoryController).Methods("POST")

	return router
}
