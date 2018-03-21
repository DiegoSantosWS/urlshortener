package main

import (
	"fmt"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/routers"
)

func init() {
	fmt.Println("Iniciando servidor...")
	err := cone.Connection()
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	fmt.Println("Server iniciado.")
}

func main() {
	routers.Routers()
}
