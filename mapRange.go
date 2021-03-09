package main

import "fmt"

func main() {

	// maps
	mmap := map[int]string{
		22: "G",
		33: "G GH",
		44: "G GH GHI",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

}
