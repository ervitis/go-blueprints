package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleMainPage)
	port := "8080"

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
