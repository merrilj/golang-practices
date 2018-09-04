package main

import "fmt"

func main() {
	s := []int{}
	for i := 0; i < 11; i++ {
		s = append(s, i)
	}

	for el := range s {
		if el % 2 == 0 {
			fmt.Println(el, "is even")
		} else {
			fmt.Println(el, "is odd")
		}
	}
}