package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Dashboard struct {
}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles("admin/views/dashboard/list/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	view.Execute(w, nil)
}
