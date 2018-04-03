package main

import (
	"fmt"
	"log"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/helpers"
	"github.com/DiegoSantosWS/encurtador-url/routers"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("Iniciando servidor...")
	err := cone.Connection()
	if err != nil {
		log.Fatal("Erro ao abrir banco de dandos: ", err.Error())
	}

	fmt.Println("Server iniciado.")

	// for example, server receive token string in request header.
	tokenstring := helpers.TokenGenerate()

	// Let's parse this by the secrete, which only server knows.
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("wsitesb"), nil
	})
	if err != nil {
		fmt.Println("ERRO MAIN, TOKEN", err.Error())
		return
	}

	if token.Valid == true {
		routers.Routers()
	}
}
