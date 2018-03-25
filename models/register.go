package models

import (
	"log"
	"net/http"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/controller"
	"github.com/DiegoSantosWS/encurtador-url/helpers"
)

//Register Abre form de registro do usuario
func Register(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"Title": "REGISTER",
	}
	if err := controller.ModelosRegister.ExecuteTemplate(w, "register.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
}

//RegisterUser registrando um usuario
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	//var dadosLogin = mux.Vars(r)
	name := r.FormValue("name")
	email := r.FormValue("email")
	user := r.FormValue("usuario")
	pass := r.FormValue("pass")

	pass, _ = helpers.HashPassword(pass)

	sql := "INSERT INTO users (nome, email, login, pass) VALUES (?, ?, ?, ?) "
	stmt, err := cone.Db.Exec(sql, name, email, user, pass)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	_, errs := stmt.RowsAffected()
	if errs != nil {
		log.Fatal(err.Error())
		return
	}
	http.Redirect(w, r, "/", 301)
}
