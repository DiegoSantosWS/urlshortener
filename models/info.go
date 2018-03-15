package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"

	"github.com/gorilla/mux"
)

//DadosLog armazena as informações em forma de json
type DadosLog struct {
	URL       string         `json:"url" db:"url"`
	IP        string         `json:"ip" db:"ip"`
	REFER     sql.NullString `json:"referencia,omitempty" db:"referencia"`
	NAVIGATOR sql.NullString `json:"browser,omitempty" db:"browser"`
	Os        sql.NullString `json:"sysoperacional,omitempty" db:"sysoperacional"`
	DATA      string         `json:"data" db:"data"`
	CONTA     string         `json:"contador" db:"contador"`
}

//Info retorna informações de acesso da url
func Info(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	tock := mux.Vars(r)
	valor := tock["id"]

	sql := "SELECT l.url, l.ip, referencia, browser, sysoperacional, DATE_FORMAT(l.data, '%d/%m/%Y %H:%i:%s') AS data, COUNT(l.id) as contador FROM logquery as l WHERE l.token = ? GROUP BY l.id ORDER BY l.id DESC"
	rows, err := cone.Db.Queryx(sql, valor)
	if err != nil {
		fmt.Println("Erro ao buscar informações de acesso. ERROR", sql, " - ", err.Error())
		return
	}

	defer rows.Close()
	dados := DadosLog{}
	var retorno []DadosLog
	for rows.Next() {
		err := rows.StructScan(&dados)
		if err != nil {
			fmt.Println("Erro ao buscar informações de acesso. ERROR ", err.Error())
			return
		}
		retorno = append(retorno, DadosLog{dados.URL, dados.IP, dados.REFER, dados.NAVIGATOR, dados.Os, dados.DATA, dados.CONTA})
	}
	retornoJSON, err := json.Marshal(retorno)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retornoJSON)
}
