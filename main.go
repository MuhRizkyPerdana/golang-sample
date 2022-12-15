package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Sample, Selamat Datang!\nApp Version: 6\nBelajar FluxCD Image Automation Update")
}

func envPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of environment variable:")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "- "+env+"\n")
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/env", envPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
