package main

import "fmt"

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) toPrint() {
	fmt.Printf("nombre %s\n", t.name)
	fmt.Printf("desc %s\n", t.description)
	fmt.Printf("completed %t\n\n", t.completed)
}

func (t *Task) setCompleted() {
	t.completed = true
}

type TaskList struct {
	tasks []*Task //que tenga ref a la memoria y no copias "*"
}

func (tl *TaskList) getTasks() []*Task {
	return tl.tasks
}

func (tl *TaskList) appedTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(id int) {
	tl.tasks = append(tl.tasks[:id], tl.tasks[id+1:]...)

	// a task desde el inicio: to id (ej 4) 0, 1, 2, 3
	//va a agregar a tasks desde ID+1 (EJ 4+1 to el final) "..." recibe todos esos parametros el append

	//final 0 1 2 3 se add 5 ... to final y se retorna ese nuevo slice

}

func main() {
	t1 := Task{
		name:        "curso go",
		description: "completar curso",
		completed:   false,
	}

	t2 := Task{
		name:        "curso html",
		description: "completar semana",
		completed:   true,
	}

	t3 := Task{
		name:        "curso css",
		description: "completar año",
		completed:   false,
	}

	t1.toPrint()
	t2.toPrint()

	list := TaskList{}
	list.appedTask(&t1) //se manda la referencia con "&"
	list.appedTask(&t2) //se manda la referencia con "&"
	list.appedTask(&t3) //se manda la referencia con "&"
	list.deleteTask(1)  // se remueve

	fmt.Println(list)
	// list.tasks[0].name = "curso golang"

	for _, task := range list.getTasks() {
		fmt.Println(task.name)
	}

	// fmt.Println(t1)

}
