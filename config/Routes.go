package config

import (
	admin "Go-Blog/admin/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMİN
	//Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)

	//SERVE FİLES

	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
