package main

import (
	"fmt"
	"net/http"
)
type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello from the Go Web Server!\tGo is soooo cool!</p1>")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:8080", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}