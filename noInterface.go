package main

import "fmt"

type perro struct {
}
type pez struct {
}
type pajaro struct {
}

func (perro) caminar() string {
	return "soy un perro y camino"
}
func (pez) nada() string {
	return "soy un pez y nado"
}
func (pajaro) vuelo() string {
	return "soy un pajaro y estoy volando"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nada())
}
func moverPajaro(p pajaro) {
	fmt.Println(p.vuelo())
}
func main() {
	p := perro{}
	pe := pez{}
	pj := pajaro{}
	moverPajaro(pj)
	moverPerro(p)
	moverPez(pe)
}
