package main

import "fmt"

type task struct {
	nombre      string
	description string
	completo    bool
}

type taskList struct {
	tasks []*task
}

func (t *task) marcarCompleta() {
	t.completo = true
}

func (t *taskList) agregarAlista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}
func (t *taskList) imprimirListaCompletados() {
	for _, task := range t.tasks {
		if task.completo {
			fmt.Println("Nombre", task.nombre)
			fmt.Println("Descripsion", task.description)
		}
	}
}
func (t *taskList) imprimirLista() {
	for _, task := range t.tasks {
		fmt.Println("Nombre", task.nombre)
		fmt.Println("Descripsion", task.description)
	}
}
func (t *task) actualizarDescripcion(descripcion string) {
	t.description = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar curso de Go",
		description: "Completar curso de Go esta semana",
	}
	t2 := &task{
		nombre:      "Completar curso de Python",
		description: "Completar curso de Python esta semana",
	}
	t3 := &task{
		nombre:      "Completar curso de Nodejs",
		description: "Completar curso de Nodejs esta semana",
	}
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarAlista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Tareas Completadas")
	lista.imprimirListaCompletados()
	mapaTareas := make(map[string]*taskList)
	mapaTareas["Jefferson"] = lista
}
