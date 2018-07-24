package main

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ROOT")
	http.Redirect(w, r, "https://ebay.com", 301)
}

func main() {
	//http.HandleFunc("/set", setHandler)
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	log.Println("Server ended")
}