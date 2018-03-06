package main

import (
	"log"
	"net/http"
)
/*
This is using the templates in the chat app in the book

type templateHandler struct {
	once 		sync.Once
	filename 	string
	templ		*template.Template
}

func (t *templateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(rw, nil)
}*/

func main() {
	http.HandleFunc("/chat", handleTemplateMainPage)
	http.HandleFunc("/", handleMainPage)
	port := "8080"

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
