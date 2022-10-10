package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Hello")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Unexpected arg", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "Name is %s", name)
}

func main() {
	fileSevrer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileSevrer)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/form", FormHandler)
	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
