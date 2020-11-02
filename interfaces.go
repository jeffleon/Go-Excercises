package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct {
}
type pez struct {
}
type pajaro struct {
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}
func (perro) mover() string {
	return "soy un perro y camino"
}
func (pez) mover() string {
	return "soy un pez y nado"
}
func (pajaro) mover() string {
	return "soy un pajaro y estoy volando"
}

func main() {
	pj := pajaro{}
	p := perro{}
	pe := pez{}
	moverAnimal(p)
	moverAnimal(pj)
	moverAnimal(pe)
}
