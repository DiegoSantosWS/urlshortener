package models

import (
	"log"
	"net/http"

	"github.com/DiegoSantosWS/encurtador-url/controller"
)

//IndexLogin carrega template de login
func IndexLogin(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "LOGIN",
	}
	if err := controller.ModelosIndex.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
}
