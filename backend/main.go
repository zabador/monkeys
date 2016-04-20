package main

import (
	"fmt"
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
	fmt.Fprintln(w, "Count: " + count + "<br />")

	matches := fmt.Sprintf("%v", t.Matches)
	fmt.Fprintln(w, "Longest: " + t.Longest + " has hit "+ matches +" time(s)<br />")

	fmt.Fprintln(w, "Duration: " + t.Duration + " Minutes<br />")
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
}
func init() {
    http.HandleFunc("/", hello)
	http.HandleFunc("/post", post)
}

