package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	print("User access / OK 200\n")
	fmt.Fprintf(w, "Golang Sample, Selamat Datang!\nApp Version: 7\nBelajar FluxCD Image Automation Update")
}

func envPage(w http.ResponseWriter, r *http.Request) {
	print("User access /env OK 200\n")
	fmt.Fprintf(w, "List of environment variable:")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "- "+env+"\n")
	}
}

func handleRequests() {
	print("Application running on port :8000\n")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/env", envPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	print("User access " + r.URL.Path + " ERROR 404 \n")
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404")
	}
}

func main() {
	handleRequests()
}
