package handlers

import "net/http"

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}
