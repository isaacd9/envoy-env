package main

import (
	"log"
	"net/http"
)

const addr = ":8000"

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})

	m.HandleFunc("/ohno", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("oops"))
	})

	middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, r.URL.String())
	})

	log.Printf("starting server on %+v", addr)
	http.ListenAndServe(addr, middleware)
}
