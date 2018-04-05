package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!_$&@()"
const (
	letterIdxBits = 6                         // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1      // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 150000000 / letterIdxBits // # of letter indices fitting in 63 bits
)

//CkTknExists Recebe o valor do token
type CkTknExists struct {
	TknExist string
}

//HashPassword encripta uma senha passada para o bando de dados
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash compara uma senha e retona verdadeiro ou false
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//CheckTokenExist verifica se o token existe se existir vamos criar outro
func CheckTokenExist(token string) bool {

	sql := "SELECT token FROM url WHERE token = ? LIMIT 1"
	rows, err := cone.Db.Queryx(sql, token)
	if err != nil {
		log.Fatal(err.Error())
	}
	var tk CkTknExists
	var id string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tk.TknExist)
		if err != nil {
			log.Fatal(err.Error())
		}
		id = tk.TknExist
	}

	switch {
	case id == "":
		return false
	case id != "":
		return true
	default:
		return false
	}
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

//Runn executa o servidor
func Runn(r *mux.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//log.Fatal(http.ListenAndServe(":3000", r))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
