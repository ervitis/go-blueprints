package main

import (
	"net/http"
	"html/template"
	"path/filepath"
)

func handleTemplateMainPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t := template.New("chat.html")
	t.ParseFiles(filepath.Join("./templates", "chat.html"))
	t.Execute(w, nil)
}

func handleMainPage(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`
		<html>
			<head>
				<title>Chat</title>
			</head>
			<body>
				Chat!<br>
				Visit the chat: <a href="/chat">Chat</a>
			</body>
		</html>
	`))
}