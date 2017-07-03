package main

import (
	"flag"
	"github.com/wouterbeets/fizzbuzz/http"
	"log"
)

func main() {
	serverAddr := ""
	flag.StringVar(
		&serverAddr,
		"address",
		"localhost:8080",
		`the addres your server run on, defaults to localhost:8080
		usuage: ./fizzbuzz -address=0.0.0.0:80`,
	)
	flag.Parse()

	s := http.NewServer(serverAddr)
	log.Fatal(s.ListenAndServe())
}
