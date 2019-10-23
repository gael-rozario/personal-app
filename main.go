package main

import (
	"net/http"
	"html/template"
)

type person struct{
	Name string
	Designation string
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
	w.WriteHeader(200)
}
func gaelpagehandler(w http.ResponseWriter, r *http.Request) {
	gael := person{Name:"Gael Rozario",Designation:"Devops Engineer"}
	html, err := template.ParseFiles("html/indextemplate.html")
	if err!= nil{
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	html.Execute(w,gael)
}

func main()  {
	
	http.HandleFunc("/", gaelpagehandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/gael", gaelpagehandler)
	http.ListenAndServe(":8080",nil)
}
