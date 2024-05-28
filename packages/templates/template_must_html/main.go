package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"PHP", 80},
		{"Python", 20},
	})
	if err != nil {
		panic(err)
	}

}
