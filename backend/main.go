package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type test_struct struct {
	Count int64
	Longest string
	Duration string
	Matches int64
}

var t test_struct

func hello(w http.ResponseWriter, r *http.Request) {
	count := fmt.Sprintf("%v", t.Count)
	io.WriteString(w, "Count: " + count + "\n")

	matches := fmt.Sprintf("%v", t.Matches)
	io.WriteString(w, "Longest: " + t.Longest + " has hit "+ matches +" time(s)\n")

	io.WriteString(w, "Duration: " + t.Duration + " Minutes\n")
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
}


func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8000", nil)
}

