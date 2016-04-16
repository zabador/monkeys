package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	flag := true
	longestMatch := make([]byte, 0, 8)
	var longestLength int


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
				if i > longestLength {
					longestLength = i
					longestMatch = s[:i+1]
					post(int64(count), longestMatch)
				}
				// if we match all of monkeys then set flag to false
				if i == len(s) - 1 {
					flag = false
				}
			} else {
				break
			}
			// check duration and post update every 5 minutes
			d := time.Now().Sub(t1)
			m := d.Minutes()
			if m > 5.0 {
				t1 = time.Now()
				fmt.Println("Total duration:", time.Now().Sub(t0))
				post(int64(count), longestMatch)

			}
		}
		count++
	}
}

func post(i int64, s []byte){
	fmt.Printf("Count: %v, Longest: %v\n", i, string(s))
}
