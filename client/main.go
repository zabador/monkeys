package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"strconv"
	"encoding/json"
	"bytes"
)

func main() {

	flag := true
	longestMatch := make([]byte, 0, 8)
	var matches int64


	s := []byte{'m','o','n','k','e','y','s'}

	count := 1


	t0 := time.Now().UTC()
	t1 := time.Now().UTC()

	for flag {
		for i := 0; i < len(s); i++ {

			// get random int
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			random := r1.Intn(123 - 97) + 97
			// check if random int equals ASCII value of letter and post new longest match
			if random == int(s[i]) {
				if i > len(longestMatch) - 1 {
					longestMatch = s[:i+1]
					matches = 1
					post(int64(count), longestMatch, matches, time.Now().Sub(t0).Minutes())
				} else if i == len(longestMatch) - 1 {
					matches++
				}

				// if we match all of monkeys then set flag to false
				if i == len(s) - 1 {
					flag = false
					dur := strconv.FormatFloat(time.Now().Sub(t0).Minutes(), 'f', -1, 64)
					fmt.Printf("WE FOUND %v in %v minutes\n", string(s), dur)
					break
				}
			} else {
				break
			}
			// check duration and post update every 5 minutes
			d := time.Now().Sub(t1)
			m := d.Minutes()
			if m > 5.0 {
				t1 = time.Now()
				post(int64(count), longestMatch, matches, time.Now().Sub(t0).Minutes())
			}
		}
		count++
	}
}

func post(i int64, s []byte, m int64, d float64){
	url := "http://monkeys-type.appspot.com/post"

	var jsonRequest struct {
		Count int64
		Longest string
		Duration string
		Matches int64
	}

	jsonRequest.Count = i
	jsonRequest.Longest = string(s)
	jsonRequest.Matches = m
	dur := strconv.FormatFloat(d, 'f', -1, 64)
	jsonRequest.Duration = dur

	j,_ := json.Marshal(jsonRequest)

	fmt.Println(string(j))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println("Failure to post")
    } else {
		resp.Body.Close()
	}

}
