package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/controller"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!_$&@()"
const (
	letterIdxBits = 6                         // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1      // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 150000000 / letterIdxBits // # of letter indices fitting in 63 bits
)

//Shorten gerador de token para url
func Shorten(w http.ResponseWriter, r *http.Request) {

	var (
		shortenURL string
		token      string
	)

	switch {
	case r.Method == http.MethodPost:

		urlOrignal := r.FormValue("shorturl")
		_, err := url.Parse(urlOrignal)
		if err != nil {
			fmt.Println("Erro url", err.Error())
			return
		}
		//tamanho := len(urlOrignal)
		tokenMD5 := GetMD5Hash(urlOrignal)
		tcknExist := RandStringBytesMaskImpr(6)

		//fmt.Println()
		if CheckTokenExist(tcknExist) == false {
			token = RandStringBytesMaskImpr(5 + 1 - 2)
		} else {
			token = RandStringBytesMaskImpr(7 * 1)
		}
		shortenURL = "http://localhost:3000/r/" + token
		_, err = InsertURL(urlOrignal, tokenMD5, token, shortenURL)
		if err != nil {
			fmt.Println("Erro: ", err.Error())
			return
		}
		data := map[string]interface{}{
			"Title": "WSITEBRASIL SHORTENING URL",
			"Url":   shortenURL,
		}
		if err := controller.ModelosHome.ExecuteTemplate(w, "home.html", data); err != nil {
			http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
			fmt.Println("Erro ao executar template", err.Error())
		}
		break
	case r.Method == http.MethodGet:
		data := map[string]interface{}{
			"Title": "WSITEBRASIL SHORTENING URL",
		}
		if err := controller.ModelosHome.ExecuteTemplate(w, "home.html", data); err != nil {
			http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
			fmt.Println("Erro ao executar template", err.Error())
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
func InsertURL(url string, tokenMD5 string, token string, shortenURL string) (string, error) {
	sql := "INSERT INTO url (url, tokenMD5, token, shortenURL) VALUES (?, ?, ?, ?) "
	stmt, err := cone.Db.Exec(sql, url, tokenMD5, token, shortenURL)
	if err != nil {
		fmt.Println("[CADEX:] Erro na inclusão do usuario", sql, " - ", err.Error())
	}

	linas, errs := stmt.RowsAffected()
	if errs != nil {
		fmt.Println("[CADEX:] Erro ao pegar numero de linhas", sql, " - ", err.Error())
	}
	//fmt.Println("Linhas: ", linas, " linas(s) afetada(s)")
	return string(linas), nil
}

//GetMD5Hash gera um token de md5
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//RandomString gera um token de string
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!_$&@()")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

//CheckTokenExist verifica se o token existe se existir vamos criar outro
func CheckTokenExist(token string) bool {
	var tokenReturned string
	if token != "" {
		sql := "SELECT token FROM url WHERE token = ? "
		linha := cone.Db.QueryRowx(sql, token)
		err := linha.Scan(&tokenReturned)
		if err != nil {
			//fmt.Println("ERROR ao buscar as informações do token", err.Error())
			return false
		}
		return true
	}
	return false
}

//RandStringBytesMaskImpr gerando token mais aleatorios
func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}