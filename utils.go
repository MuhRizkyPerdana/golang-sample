package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func stringGen(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyz")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func numberGen(n int) string {
	var chars = []rune("0123456789")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func timeTrack(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed
}

func sendtoTempo(body []byte, tempoURL string, traceId string, ismain bool) {
	time.Sleep(3 * time.Second)
	url := tempoURL

	var jsonStr = body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		create_log(err.Error())
	}
	defer resp.Body.Close()

	if ismain {
		create_log("trace pushed " + tempoURL + " ID " + traceId)
	}

}

func create_log(msg string) {
	if len(os.Args) > 1 {
		if logFile {
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
			fmt.Println(msg)
		}
	} else {
		fmt.Println(msg)
	}
}
