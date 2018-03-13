package routers

import (
	"net/http"

	"github.com/DiegoSantosWS/encurtador-url/models"
	"github.com/gorilla/mux"
)

func Routers() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", models.Shorten)
	r.HandleFunc("/encurt-url", models.Shorten)
	r.HandleFunc("/r/{token}", models.Redirection)
	r.HandleFunc("/list", models.ListResults)
	r.HandleFunc("/info/{id}", models.Info)

	http.ListenAndServe(":3000", r)
}
