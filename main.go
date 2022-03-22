package main

import (
	admin_models "Go-Blog/admin/models"
	"Go-Blog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	post := admin_models.Post{}.Get(5)
	post.Delete()
	http.ListenAndServe(":8080", config.Routes())
}
