package main

import "net/http"

func handleMainPage(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`
		<html>
			<head>
				<title>Chat</title>
			</head>
			<body>
				Chat!
			</body>
		</html>
	`))
}