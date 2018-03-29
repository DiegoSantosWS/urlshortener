package models

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/gorilla/mux"
)

//User seta id se usuario ja estiver cadastrado
type User struct {
	ID int
}

//CkTkn seta id se token ja estiver cadastrado
type CkTkn struct {
	ID int
}

type RetornoALTER struct {
	NewToken  string `json:"newToken"`
	Msg       string `json:"msg"`
	Verifcado bool
}

//CheckCad verifica se usuario ja está cadastrado via o email
func CheckCad(w http.ResponseWriter, r *http.Request) {
	recept := mux.Vars(r)
	email := recept["email"]

	sql := "SELECT id FROM users WHERE email = ? LIMIT 1"
	rows, err := cone.Db.Queryx(sql, email)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var user User
	var retornoJ int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		retornoJ = user.ID
	}
	retornoJSON, err := json.Marshal(retornoJ)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retornoJSON)
}

//CheckToken verifica se token existe se não existe realiza a alteração
func CheckToken(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)

	recept := mux.Vars(r)
	newToken := recept["newToken"]
	token := recept["tkn"]

	sql := "SELECT id FROM url WHERE token = ? LIMIT 1"
	rows, err := cone.Db.Queryx(sql, newToken)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var tk CkTkn
	var id int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tk.ID)
		if err != nil {
			log.Fatal(err.Error())
		}
		id = tk.ID
	}
	var retornoJS []RetornoALTER
	m := RetornoALTER{}
	cod := strconv.Itoa(id)
	if cod == "0" {
		sql := "UPDATE url SET `token` = ? WHERE token = ?"
		stmt, err := cone.Db.Exec(sql, newToken, token)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, errs := stmt.RowsAffected()
		if errs != nil {
			log.Fatal(errs.Error())
		}
		m.NewToken = newToken
		m.Msg = "Token altered"
		m.Verifcado = true
		retornoJS = append(retornoJS, RetornoALTER{m.NewToken, m.Msg, m.Verifcado})
	} else {
		m.NewToken = newToken
		m.Msg = "Unchanged, possible token already exists!"
		m.Verifcado = false
		retornoJS = append(retornoJS, RetornoALTER{m.NewToken, m.Msg, m.Verifcado})
	}

	retornoJSON, err := json.Marshal(retornoJS)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retornoJSON)
}
