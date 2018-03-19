package models

import (
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
