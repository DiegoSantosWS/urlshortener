package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/controller"
)

//OriginalURL recebe o valor do banco em forma de struct
type OriginalURL struct {
	URL string `json:"url" db:"url"`
}

//BrowserReferer retornar os dados do browswer
type BrowserReferer struct {
	Browser string `json:"browser" db:"browser"`
	Click   string `json:"clicks"`
}

//AnalyticsResults retorna os dados para template
func AnalyticsResults(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	var cod = mux.Vars(r)
	id := cod["id"]

	totalClcks := CountClicks(id)

	sql := "SELECT url FROM url WHERE token = ? LIMIT 1"
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("[analytics] erro ao buscar url orginal: ", sql, " - ", err.Error())
		return
	}

	defer rows.Close()
	dadosURL := OriginalURL{}
	var original string
	for rows.Next() {
		err := rows.StructScan(&dadosURL)
		if err != nil {
			fmt.Println("Erro ao renderizar a url para variavel", err.Error())
			return
		}
		original = dadosURL.URL
	}

	data := map[string]interface{}{
		"SubTitle":       "Analytics data for",
		"Short":          "https://wsib.ws:3000/" + id,
		"Original":       original,
		"tokenAnalytcis": id,
		"TotalClicks":    totalClcks,
	}
	if err := controller.ModelosAnalytics.ExecuteTemplate(w, "analytics-wd.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
		fmt.Println("Erro ao executar template", err.Error())
	}
}

//AnalytcsChart Montando grafico
func AnalytcsChart(w http.ResponseWriter, r *http.Request) {

	cod := mux.Vars(r)
	var id = cod["id"]
	totalClcks := CountClicks(id)

	sql := "SELECT url FROM url WHERE token = ? LIMIT 1"
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("[analytics] erro ao buscar url orginal: ", sql, " - ", err.Error())
		return
	}

	defer rows.Close()
	dadosURL := OriginalURL{}
	var original string
	for rows.Next() {
		err := rows.StructScan(&dadosURL)
		if err != nil {
			fmt.Println("Erro ao renderizar a url para variavel", err.Error())
			return
		}
		original = dadosURL.URL
	}

	data := map[string]interface{}{
		"SubTitle":       "Analytics data for",
		"Short":          "http://wsib.ws:3000/" + id,
		"Original":       original,
		"tokenAnalytcis": id,
		"TotalClicks":    totalClcks,
	}
	fmt.Println(id)
	if err := controller.ModelosAnalyticsChart.ExecuteTemplate(w, "chart.html", data); err != nil {
		http.Error(w, "[CHART] Erro in the execute template", http.StatusInternalServerError)
		fmt.Println("Erro ao executar template", err.Error())
	}
}

//GetBrowsersReferer retornar a referencia dos browsers
func GetBrowsersReferer(w http.ResponseWriter, r *http.Request) {
	cod := mux.Vars(r)
	var id = cod["id"]
	//totalClcks := CountClicks(id)

	sql := "SELECT DISTINCT(browser) as browser FROM logquery WHERE token = ? "
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("[analytics chart] erro ao buscar url orginal: ", sql, " - ", err.Error())
		return
	}

	defer rows.Close()
	dadosB := BrowserReferer{}
	var browser []BrowserReferer
	for rows.Next() {
		err := rows.StructScan(&dadosB)
		if err != nil {
			fmt.Println("Erro ao renderizar a url para variavel", err.Error())
			return
		}
		dadosB.Click = "2"
		teste := dadosB.Click
		browser = append(browser, BrowserReferer{dadosB.Browser, teste})
	}
	retornoJSON, err := json.Marshal(browser)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retornoJSON)
}
