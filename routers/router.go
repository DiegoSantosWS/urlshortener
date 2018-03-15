package routers

import (
	"log"
	"net/http"

	"github.com/DiegoSantosWS/encurtador-url/models"
	"github.com/gorilla/mux"
)

//Routers instacia as rotas do sistema
func Routers() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", models.IndexLogin).Methods("GET")
	r.HandleFunc("/register", models.Register).Methods("GET")
	r.HandleFunc("/register-user", models.RegisterUser)
	r.HandleFunc("/login", models.Login)
	r.HandleFunc("/home", models.Shorten)
	r.HandleFunc("/encurt-url", models.Shorten)
	r.HandleFunc("/r/{token}", models.Redirection)
	r.HandleFunc("/list", models.ListResults)
	r.HandleFunc("/analytics-wd/{id}", models.AnalyticsResults)
	r.HandleFunc("/info/{id}", models.Info)
	r.HandleFunc("/check-cad/{email}", models.CheckCad)

	log.Fatal(http.ListenAndServe(":3000", r))
}
