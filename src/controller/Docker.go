package controller

import "fmt"
import "html/template"
import "net/http"
import . "model"

func Docker(response http.ResponseWriter, request *http.Request) {
	liumiaocn := Person{Id: 1001, Name: "liumiaocn", Country: "China"}

	tmpl, err := template.ParseFiles("./view/docker.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response, liumiaocn)
}
