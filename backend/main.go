package main

import (
	"flag"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func main() {
	ports := flag.String("ports", ":8080", "Ports to run server to")
	flag.Parse()
	h := &Handler{}
	http.HandleFunc("/", h.HelloHandler)
	if err := http.ListenAndServe(*ports, nil); err != nil {
		log.Fatal(err)
	}
}
