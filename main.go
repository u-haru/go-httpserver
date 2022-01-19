package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	localHost string
	localDir  string
)

func main() {
	flag.StringVar(&localHost, "p", "0.0.0.0:8080", "Proxy:port")
	flag.StringVar(&localDir, "d", ".", "/path/to/dir")
	flag.Parse()

	log.Println("Start serving on", localHost)
	http.Handle("/", http.FileServer(http.Dir(localDir)))
	if err := http.ListenAndServe(localHost, nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
