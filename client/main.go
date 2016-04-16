package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	flag := true
	longest := make([]byte, 0, 8)

	s := []byte{'m','o','n','k','e','y','s'}

	count := 1

	for flag {
		for i := 0; i < len(s); i++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			random := r1.Intn(123 - 97) + 97
			if random == int(s[i]) {
				if i > len(longest) {
					longest = s[:i+1]
					post(int64(count), longest)
				}
				if i == len(s) - 1 {
					flag = false
				}
			} else {
				break
			}
		}
		count++
	}
}

func post(i int64, s []byte){
	fmt.Printf("Count: %v, Longest: %v\n", i, string(s))
}
