package config

import (
	admin "Go-Blog/admin/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMİN
	r.GET("/admin", admin.Dashboard{}.Index)

	//SERVE FİLES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
