package models

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/controller"
	"github.com/DiegoSantosWS/encurtador-url/helpers"
)

//Shorten gerador de token para url
func Shorten(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)

	var (
		shortenURL string
		token      string
		persona    string
	)

	switch {
	case r.Method == http.MethodPost:

		urlOrignal := r.FormValue("shorturl")
		_, err := url.Parse(urlOrignal)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		//tamanho := len(urlOrignal)
		tokenMD5 := helpers.GetMD5Hash(urlOrignal)
		tcknExist := helpers.RandStringBytesMaskImpr(4)

		fmt.Println("Existe: ", tcknExist)
		if helpers.CheckTokenExist(tcknExist) == true {
			sid, _ := helpers.New(1, helpers.DefaultABC, 2)
			tokenGenerate, _ := sid.Generate()
			token = tokenGenerate
		} else {
			token = tcknExist
		}
		fmt.Println("Passou\n", token)
		shortenURL = r.Host + "/" + token
		session, _ := Store.Get(r, "logado")
		company := session.Values["ID"]
		_, err = InsertURL(urlOrignal, tokenMD5, token, shortenURL, persona, company)
		if err != nil {
			log.Fatal(err.Error())
		}
		data := map[string]interface{}{
			"Title": "WSITEBRASIL SHORTENING URL",
			"Url":   shortenURL,
		}
		if err := controller.ModelosHome.ExecuteTemplate(w, "home.html", data); err != nil {
			http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
			log.Fatal(err.Error())
		}
		break
	case r.Method == http.MethodGet:
		data := map[string]interface{}{
			"Title": "WSITEBRASIL SHORTENING URL",
		}
		if err := controller.ModelosHome.ExecuteTemplate(w, "home.html", data); err != nil {
			http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
			log.Fatal(err.Error())
		}
		break
	default:
		break
	}
}

/*InsertURL inserindo a url no banco para salvar as informações
 * url recebe uma string,
 * tokenMD5 recebe uma string,
 * token recebe uma string,
 * shortenURL  recebe uma string
 * @return string or error
 */
func InsertURL(url string, tokenMD5 string, token string, shortenURL string, persona string, company interface{}) (string, error) {
	sql := "INSERT INTO url (url, tokenMD5, token, shortenURL, personalised, company) VALUES (?, ?, ?, ?, ?, ?) "
	stmt, err := cone.Db.Exec(sql, url, tokenMD5, token, shortenURL, persona, company)
	if err != nil {
		fmt.Println("[CADEX:] Erro na inclusão do usuario", sql, " - ", err.Error())
	}

	linas, errs := stmt.RowsAffected()
	if errs != nil {
		log.Fatal(errs.Error())
	}

	return string(linas), nil
}
