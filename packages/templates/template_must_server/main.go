package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Courses{
			{"Go", 40},
			{"PHP", 80},
			{"Python", 20},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)

}
