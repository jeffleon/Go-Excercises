package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operator(entrada string, operador string) int {
	valores := strings.Split(entrada, operador)
	operador1 := parseo(valores[0])
	operador2 := parseo(valores[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 * operador2
	default:
		fmt.Println(operador + " no es una operacion valida")
		return 0
	}
}
func parseo(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}
func inputScan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func main() {
	calculadora := calc{}
	entrada := inputScan()
	operador := inputScan()
	resultado := calculadora.operator(entrada, operador)
	fmt.Println(resultado)
}
