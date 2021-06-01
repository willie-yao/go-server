package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	server := os.Getenv("SERVER_ADDR")
	resp, err := http.Get(server)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
