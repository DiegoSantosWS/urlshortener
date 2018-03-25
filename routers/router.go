package routers

import (
	"fmt"
	"log"
	"net"
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
	r.HandleFunc("/logout", models.Logout)
	r.HandleFunc("/home", models.Shorten)
	r.HandleFunc("/encurt-url", models.Shorten)
	r.HandleFunc("/{token}", models.Redirection).Methods("GET")
	r.HandleFunc("/analytics-wd/{id}", models.AnalyticsResults)
	r.HandleFunc("/check-cad/{email}", models.CheckCad)

	r.HandleFunc("/list/", models.ListResults)
	r.HandleFunc("/info/{id}", models.Info)
	r.HandleFunc("/analyticsChar/{id}", models.AnalytcsChart)
	r.HandleFunc("/info-browser/{id}", models.GetBrowsersReferer)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println("[HELPERS-PORT] ERRO AO RECUPERAR A PORTA", err.Error())
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("Seu sistema está rodando em: http://localhost:%d\n\r", port)
	log.Fatal(http.Serve(listener, r))

}
