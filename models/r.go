package models

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/gorilla/mux"
	"github.com/ua-parser/uap-go/uaparser"
)

//DadosUsuarios armazena os dados do login recebido do banco...
type DadosUsuarios struct {
	URL string
}

//Redirection realiza o redirecionamento para url correta
func Redirection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenURL := vars["token"]
	referencia := r.Header.Get("Referer")
	//browser := r.Header.Get("User-Agent")
	uagent := r.Header.Get("User-Agent")
	parser, err := uaparser.New("regexes.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	client := parser.Parse(uagent)
	browser := client.UserAgent.Family
	sysoperacional := client.Os.Family

	if tokenURL == "" {
		http.Error(w, "Não foi enviado token válido. Verifique.", http.StatusBadRequest)
		fmt.Println("Erro ao encontrar o token ")
		return
	}
	sql := "SELECT url FROM url WHERE token = ? "
	linha, err := cone.Db.Queryx(sql, tokenURL)
	if err != nil {
		http.Error(w, "[ERRO] Código inixistente. verifique.", http.StatusInternalServerError)
		fmt.Println("[ERRO] Usuário não encontrado", err.Error())
		return
	}
	defer linha.Close()
	u := DadosUsuarios{}
	for linha.Next() {
		err := linha.Scan(&u.URL)
		if err != nil {
			http.Error(w, "[ERRO] Usuário não encontrado", http.StatusInternalServerError)
			fmt.Println("[ERRO] Usuário não encontrado", err.Error())
			return
		}
	}

	if err := InsertClick(u.URL, tokenURL, referencia, browser, sysoperacional, w, r); err == false {
		fmt.Println("Não foi possível inserir as informações")
	}

	http.Redirect(w, r, u.URL, 302)
}

//InsertClick gera log de clicks
func InsertClick(url, token, referencia, browser, sysoperacional string, w http.ResponseWriter, r *http.Request) bool {

	ip, err := getIPAdress(w, r)
	if err != nil {
		fmt.Println(err.Error())
	}

	sql := "insert into logquery (url, token, ip, data, referencia, browser, sysoperacional) values (?,?,?,?,?,?,?)"
	_, err = cone.Db.Exec(sql, url, token, ip, time.Now(), referencia, browser, sysoperacional)
	if err != nil {
		fmt.Println("erro: ", err.Error())
		return false
	}
	return true
}

//getIPAdress retorna edenço de ip
func getIPAdress(w http.ResponseWriter, r *http.Request) (string, error) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	userip := ip
	return userip, nil
}
