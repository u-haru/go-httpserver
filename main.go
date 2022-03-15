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
	flag.StringVar(&localHost, "p", "0.0.0.0:8080", "Address:Port")
	flag.StringVar(&localDir, "d", ".", "/path/to/dir")
	flag.Parse()

	log.Println("Start serving on", localHost)
	handler := http.FileServer(http.Dir(localDir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store") //last-modified等で確認取れない限り再取得
		handler.ServeHTTP(w, r)
	})
	if err := http.ListenAndServe(localHost, nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
