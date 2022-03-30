package main

import (
	admin_models "Go-Blog/admin/models"
	"Go-Blog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
