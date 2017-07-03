package main

import (
	"fizzbuzz/http"
	"flag"
	"log"
)

func main() {
	serverAddr := ""
	flag.StringVar(
		&serverAddr,
		"address",
		"localhost:8000",
		`the addres your server run on, defaults to localhost:8000
		usuage: ./fizzbuzz -address=0.0.0.0:80`,
	)
	flag.Parse()

	s := http.NewServer(serverAddr)
	log.Fatal(s.ListenAndServe())
}
