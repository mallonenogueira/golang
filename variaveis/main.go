package main

import (
	"fmt"
	"reflect"
)

func main() {
	variaveisAula1()

	fmt.Println("---------------")

	inferenciaAula3()
}

func variaveisAula1() {
	var nome string = "Mallone"
	fmt.Println("Olá: ", nome)

	var idade int = 31
	fmt.Println("Idade: ", idade)

	var altura float32 = 1.78
	fmt.Println("Altura: ", altura)

	var testeSemIniciar int
	var testeSemIniciarString string
	fmt.Println("Teste: ", testeSemIniciar, " String: ", testeSemIniciarString)
}

func inferenciaAula3() {
	var teste = "Mallone"

	fmt.Println(teste)

	declaracaoCurta := 1.8

	fmt.Println(reflect.TypeOf(declaracaoCurta))

}
