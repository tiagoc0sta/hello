package main

/*--------------------RETORNO NORMAL-----------------------------------------

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	res, err:= soma(10,10)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}

func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}*/

/*----------------------RETORNO MULTIPLO------------------------------------------

import "fmt"


func main() {
	resultado, str := soma(10, 20)

	fmt.Println(resultado, str)
}

func soma(a int, b int) (int, string) {
	return a + b, "somou!"
}*/


import "fmt"


func main() {
	//resultado := soma(10, 20)
	resultado :=somaTudo(3,5,7,8)

	fmt.Println(resultado)
}

func soma(a int, b int) (result int) { //retorno nomeado
	result = a + b
	return 
}


func somaTudo(x...int) int{
	resultado :=0

	for _, v :+ range x {
		resultado += v
	}

	return resultado
}
