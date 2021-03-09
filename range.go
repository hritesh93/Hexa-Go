package main

import "fmt"

func main() {
	rvar := []string{"A", "A ABC", "A ABC ABCABC"}

	for i, j := range rvar {
		fmt.Println(i, j)
	}
}
