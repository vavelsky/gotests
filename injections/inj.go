package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie")
}

// Greet welcomes
func Greet(writer io.Writer, name string) {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingHandler))
}

// MyGreetingHandler says helo over http
func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
