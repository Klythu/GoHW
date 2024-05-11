package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyaddr string = "localhost:8080"

var (
	count      int    = 0
	firstport  string = "http://localhost:8082"
	secondport string = "http://localhost:8081"
)

func proxyhand(w http.ResponseWriter, r *http.Request) {
	textByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if count == 0 {
		_, err := http.Post(firstport, "application/json; charset=utf-8", bytes.NewBuffer(textByte))
		if err != nil {
			log.Fatalln(err)
		}

		count++
		return
	}
	if count == 1 {
		_, err := http.Post(secondport, "application/json; charset=utf-8", bytes.NewBuffer(textByte))
		if err != nil {
			log.Fatalln(err)
		}
		count--
		return
	}
}
func main() {
	http.HandleFunc("/", proxyhand)
	log.Fatalln(http.ListenAndServe(proxyaddr, nil))
}
