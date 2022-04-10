package main

import "fmt"

//las estructuras se puede relacionar

type User struct {
	nombre string
	email  string
	status bool
}

type Student struct {
	user User //la relacion variable de tipo usuario diferente de la herencia
	code string
}

//1:N
type Curso struct {
	title  string
	videos []Videos // un curso tiene N videos
}

//1:1
type Videos struct {
	title string
	curso Curso //un video tiene 1 curso
}

func main() {
	//relacion de 1:N

	video1 := Videos{
		title: "intro",
	}
	video2 := Videos{
		title: "end",
	}
	curso1 := Curso{
		title:  "curso go",
		videos: []Videos{video1, video2},
	}

	/* 	vl := make([]Videos, 2)
	   	vl = append(vl, video1)
	   	vl = append(vl, video2)

	   	curso2 := Curso{
	   		title:  "curso c",
	   		videos: vl,
	   	} */

	video1.curso = curso1
	video2.curso = curso1

	// fmt.Println(video1.curso.title)

	for _, video := range curso1.videos {
		fmt.Println(video.title)
	}
}
