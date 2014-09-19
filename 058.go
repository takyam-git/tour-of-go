package main

import (
	"net/http"
)

type String string

func (s String) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(s))
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(s.Greeting + s.Punct + s.Who))
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
