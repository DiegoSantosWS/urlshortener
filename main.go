package main

import (
	"fmt"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/routers"
)

func init() {
	err := cone.Connection()
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
}

func main() {
	routers.Routers()
}
