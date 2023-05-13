package playground

import (
	"log"
	"net/http"
	"testing"
)

type server int

// still interface function
func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello world"))
}

func otherName(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello world"))
}

type server2 int

// still interface function
func (h server2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello world"))
}

func Test1(t *testing.T) {
	var s server

	http.ListenAndServe("localhost:8080", &s)

	var s2 server2
	http.ListenAndServe("localhost:8080", s2)

	http.ListenAndServe("localhost:8080", http.HandlerFunc(otherName))
}
