package models

import (
	"fmt"
	"log"
	"net/http"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/helpers"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("1234567890dhskai#sobn")
	Store = sessions.NewCookieStore(key)
)

//UsuarioLogin struct para armazenar os dados do login
type UsuarioLogin struct {
	User string
	Pass string
}

//UsuarioLogado armazena os dados do login recebido do banco...
type UsuarioLogado struct {
	ID      int    `db:"id"`
	Nome    string `db:"nome"`
	Email   string `db:"email"`
	Usuario string `db:"login"`
	Senha   string `db:"pass"`
}

//Logout faz logout com usuario
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "logado")
	session.Values["autorizado"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

//Login carrega uma template
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		if r.FormValue("username") == "" && r.FormValue("password") == "" {
			return
		}
		//Vamos receber os dados do formulario
		usr := UsuarioLogin{
			User: r.FormValue("username"),
			Pass: r.FormValue("password"),
		}
		//Vamos verificar se usuario se senha Existe
		user := UsuarioLogado{}
		sql := "SELECT id, nome, email, login, pass FROM users WHERE login = ? "
		rows, err := cone.Db.Queryx(sql, usr.User)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		defer rows.Close()
		//for para buscar os dados do usuario
		for rows.Next() {
			err := rows.StructScan(&user)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
		}
		entrar := helpers.CheckPasswordHash(usr.Pass, user.Senha)
		if entrar != false {
			session, _ := Store.Get(r, "logado")
			session.Values["ID"] = user.ID
			session.Values["Nome"] = user.Nome
			session.Values["Email"] = user.Email
			session.Values["autorizado"] = true
			session.Save(r, w)
			CheckSession(w, r)
			http.Redirect(w, r, "/home", 301)
		}
		//LogAcesso(user.Nome, user.Tipo, "Falha")
		http.Redirect(w, r, "/", 301)
	}
}

//CheckSession verifica uma sessão
func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "logado")
	fmt.Println(session.Values["autorizado"])
	if auth, ok := session.Values["autorizado"].(bool); !ok || !auth {
		if auth == false {
			http.Redirect(w, r, "/", 301)
			//http.Error(w, "Acesso negado", http.StatusInternalServerError)
		}
		http.Error(w, "Acesso negado", http.StatusInternalServerError)
		return
	}
}
