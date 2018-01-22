package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World\n")
}

func jobsHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Jobs\n")
}

func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/jobs", jobsHandler)

	http.ListenAndServe(":8080", nil)
}