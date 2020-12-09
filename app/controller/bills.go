package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/iavealokin/MDSite/app/model"

	"github.com/julienschmidt/httprouter"
)

func GetBillUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bills, err := model.GetBillUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	main := filepath.Join("public", "html", "index.html")
	common := filepath.Join("public", "html", "template.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.ExecuteTemplate(rw, "bills", bills)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}
