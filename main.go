package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	create_log("User access / OK 200")
	fmt.Fprintf(w, "Golang Sample, Selamat Datang!\nApp Version: 13")
}

func envPage(w http.ResponseWriter, r *http.Request) {
	create_log("User access /env OK 200")
	fmt.Fprintf(w, "List of environment variable:")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "- "+env+"\n")
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	create_log("User access /api OK 200")
	if !trace {
		fmt.Fprintf(w, string("Tracing not enabled"))
		create_log("User access /api ERROR 404 Tracing not enabled")
		return
	}

	var c1, _ = strconv.Atoi(r.URL.Query().Get("data-service"))
	var c2, _ = strconv.Atoi(r.URL.Query().Get("product-service"))
	var c3, _ = strconv.Atoi(r.URL.Query().Get("payment-service"))
	if c1 == 0 {
		c1 = 2
	}
	if c2 == 0 {
		c2 = 3
	}
	if c3 == 0 {
		c3 = 4
	}

	start := time.Now()
	var timeStampMicro = time.Now().UnixNano() / int64(time.Microsecond)
	var id = numberGen(5)
	var trcId = numberGen(7)

	traceChildOne := apiTraceChild(id, trcId, c1)
	traceChildTwo := apiTraceChild(id, trcId, c2)
	traceChildThree := apiTraceChild(id, trcId, c3)

	var trace = []TraceTempo{
		{
			ID:        id,
			TraceID:   trcId,
			Duration:  timeTrack(start).Microseconds(),
			Name:      os.Args[0],
			Timestamp: timeStampMicro,
			Tags: Tags{
				HttpMethod: "GET",
				HttpPath:   "/api",
			},
			LocalEndpoint: LocalEndpoint{
				ServiceName: "goapp",
			},
		},
	}

	resTrace, _ := json.Marshal(trace)
	resTraceChildOne, _ := json.Marshal(traceChildOne)
	resTraceChildTwo, _ := json.Marshal(traceChildTwo)
	resTraceChildThree, _ := json.Marshal(traceChildThree)
	sendtoTempo(resTrace, traceURL, trace[0].TraceID, true)
	sendtoTempo(resTraceChildOne, traceURL, traceChildOne[0].TraceID, false)
	sendtoTempo(resTraceChildTwo, traceURL, traceChildTwo[0].TraceID, false)
	sendtoTempo(resTraceChildThree, traceURL, traceChildThree[0].TraceID, false)

	fmt.Fprintf(w, string(resTrace))

}

func apiTraceChild(parentId string, traceId string, tms int) []TraceTempoChild {
	create_log("Call function [apiTraceChild] OK")
	start := time.Now()
	var timeStampMicro = time.Now().UnixNano() / int64(time.Microsecond)

	time.Sleep(time.Second * time.Duration(tms))

	var traceChild = []TraceTempoChild{
		{
			ID:        numberGen(5),
			TraceID:   traceId,
			Duration:  timeTrack(start).Microseconds(),
			Name:      os.Args[0],
			Timestamp: timeStampMicro,
			ParentId:  parentId,
			LocalEndpoint: LocalEndpoint{
				ServiceName: "goapp",
			},
		},
	}

	return traceChild

}

func handleRequests() {
	create_log("Application running on port :8000")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/env", envPage)
	http.HandleFunc("/api", api)
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

func main() {
	for _, value := range os.Args {
		if value == "--log-file" {
			logFile = true
		}
		if strings.Split(value, "=")[0] == "--trace-url" {
			trace = true
			traceURL = strings.Split(value, "=")[1]
		}
	}
	handleRequests()
}

var logFile = false
var trace = false
var traceURL = ""
