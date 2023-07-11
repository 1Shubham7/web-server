package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")

	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "did not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported mf", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello Ji")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err !=nil {
		fmt.Fprint(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprint(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s", address)
}