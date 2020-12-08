package controller

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

//StartPage for users ...
func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//указываем путь к нужному файлу
	path := filepath.Join("public", "html", "login.html")
	common := filepath.Join("public", "html", "common.html")
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
