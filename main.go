package main

import (
	"Go-Blog/config"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", config.Routes())
}
