package main

import (
	"flag"
	"fmt"
	"github.com/dnagikh/website/cmd/web/router"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 8006, "specify port number")
)

func main() {
	flag.Parse()

	mux := router.NewRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	log.Printf("Running server on http://localhost%s\n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
