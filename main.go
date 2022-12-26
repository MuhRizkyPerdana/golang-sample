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
	create_log("User access / OK 200")
	fmt.Fprintf(w, "Golang Sample, Selamat Datang!\nApp Version: 10")
}

func envPage(w http.ResponseWriter, r *http.Request) {
	create_log("User access /env OK 200")
	fmt.Fprintf(w, "List of environment variable:")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "- "+env+"\n")
	}
}

func handleRequests() {
	create_log("Application running on port :8000")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/env", envPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	create_log("User access " + r.URL.Path + " ERROR 404")
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404")
	}
}

func Hello(name string) string {
	return "Hi, " + name
}

func create_log(msg string) {
	if len(os.Args) > 1 {
		if os.Args[1] == "--log-file" {
			f, err := os.OpenFile(os.Args[0]+".log",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(msg + "\n"); err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println(msg + "\n")
		}
	} else {
		fmt.Println(msg)
	}
}

func main() {
	handleRequests()
}
