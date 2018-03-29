package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
)

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

func Runn(r *mux.Router) {
	/*
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer listener.Close()

		port := listener.Addr().(*net.TCPAddr).Port
	*/
	port := os.Getenv("PORT")
	//fmt.Print(port)
	//log.Fatal(http.ListenAndServe(":3000", r))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
