package main

import (
	"fmt"
	"net/http"
)

func main() {

	//setting up routes
	http.HandleFunc("/", publicResources)
	http.HandleFunc("/secure", protectedResources)

	//starting the server
	fmt.Println("Server is running on port: 8080")
	http.ListenAndServe(":8080", nil)

}

func publicResources(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/public.html")
}

func protectedResources(w http.ResponseWriter, r *http.Request) {
	username := "admin"
	password := "admin"

	//extracting uname & passwrd from auth request
	user, pass, ok := r.BasicAuth()

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	if ok {
		if username != user || password != pass {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
	http.ServeFile(w, r, "static/private.html")
}
