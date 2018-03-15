package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/gorilla/mux"
)

//User seta id se usuario ja estiver cadastrado
type User struct {
	ID int
}

//CheckCad verifica se usuario ja está cadastrado via o email
func CheckCad(w http.ResponseWriter, r *http.Request) {
	recept := mux.Vars(r)
	email := recept["email"]

	sql := "SELECT id FROM users WHERE email = ? LIMIT 1"
	rows, err := cone.Db.Queryx(sql, email)
	if err != nil {
		fmt.Println("[CHECK] Erro ao executar consulta", sql, " - ", err.Error())
		return
	}

	var user User
	var retornoJ int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID)
		if err != nil {
			fmt.Println("[CHECK] Erro ao executar ", sql, " - ", err.Error())
			return
		}
		retornoJ = user.ID
	}
	retornoJSON, err := json.Marshal(retornoJ)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retornoJSON)
}
