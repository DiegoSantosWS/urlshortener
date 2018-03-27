package models

import (
	"encoding/json"
	"log"
	"net/http"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
)

//ListaURL retorna os dados da url
type ListaURL struct {
	URL       string `json:"url" db:"url"`
	Tkn       string `json:"token" db:"token"`
	Count     string `json:"total"`
	ShrtenURL string `json:"shortenURL" db:"shortenURL"`
}

//ListResults lista todos os resultados cadastrados
func ListResults(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	//recebendo a sessão para buscar url's por empresa/usuario
	session, _ := Store.Get(r, "logado")
	id := session.Values["ID"]

	sql := "SELECT u.url, u.token, u.shortenURL FROM url as u WHERE u.company = ? ORDER BY u.id DESC "
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer rows.Close()
	//dadosB := BrowserReferer{}
	//var browser []BrowserReferer
	dadosURL := ListaURL{}
	var listURL []ListaURL
	for rows.Next() {

		err := rows.StructScan(&dadosURL)
		if err != nil {
			log.Fatal(err.Error())
		}
		dadosURL.Count = CountClicks(dadosURL.Tkn)
		listURL = append(listURL, ListaURL{dadosURL.URL, dadosURL.Tkn, dadosURL.Count, dadosURL.ShrtenURL})
	}
	listURLData, err := json.Marshal(listURL)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(listURLData)
}

//CountClicks conta total de clicks que token obteve
func CountClicks(token string) string {

	sql := "SELECT count(*) as total  FROM logquery WHERE token = ? "
	rows, err := cone.Db.Queryx(sql, token)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	defer rows.Close()

	var total string
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			log.Fatal(err.Error())
			return ""
		}
	}
	return string(total)
}
