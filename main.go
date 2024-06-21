package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./client"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler) 
	http.HandleFunc("/hello", helloHandler)
	
	fmt.Printf("Server is statring at Port 8080!\n")

	if err := http.ListenAndServe(":8080", nil); 
	err != nil {
		log.Fatal(err);
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "Page Not Found!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not accepted!", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello from Golang Server!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Err: %v", err)
		return
	}
	fmt.Fprint(w, "Post requested sent successfully!\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name: %s\n", name);
	fmt.Fprintf(w, "Email: %s\n", email);
}