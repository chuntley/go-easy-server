package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	directory := flag.String("directory", ".", "The directory to serve")
	serverPort := flag.String("port", "4000", "The server port")
	flag.Parse()

	fmt.Printf("Running server on port %s, serving files from %s", *serverPort, *directory)
	panic(http.ListenAndServe(":"+*serverPort, http.FileServer(http.Dir(*directory))))
}
