package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
)

//ListResults lista todos os resultados cadastrados
func ListResults(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT u.url, u.token  FROM url as u ORDER BY u.id DESC "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", sql, " - ", err.Error())
		return
	}

	type Groups struct {
		URL   string `json:"url"`
		Tkn   string `json:"token"`
		Count string `json:"total"`
	}
	defer rows.Close()

	var groups []Groups
	for rows.Next() {
		var URL string
		var Tkn string
		var Cont string

		rows.Scan(&URL, &Tkn)
		Cont = CountClicks(Tkn)
		groups = append(groups, Groups{URL, Tkn, Cont})
	}
	groupData, err := json.Marshal(groups)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(groupData)
}

//CountClicks conta total de clicks que token obteve
func CountClicks(token string) string {
	sql := "SELECT count(*) as total  FROM logquery WHERE token = ? "
	rows, err := cone.Db.Queryx(sql, token)
	if err != nil {
		fmt.Println("[CONTADOR] Erro ao buscar informações de GRUPO: ", sql, " - ", err.Error())
		return ""
	}

	defer rows.Close()

	var total string
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			fmt.Println("[CONTADOR] Erro ao buscar total de clickes: - ", err.Error())
			return ""
		}
	}
	return string(total)
}
