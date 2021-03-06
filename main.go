package main

import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	fmt.Fprintf(w, "First Name = %s\n", firstName)
	fmt.Fprintf(w, "Last Name = %s\n", lastName)
}

func functionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/function" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "This is not a supported METHOD of this function", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "this is a function print statement")
}

func main() {
	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/function", functionHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}