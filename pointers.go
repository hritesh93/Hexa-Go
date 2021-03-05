package main

import "fmt"

func main() {
	star := "Polaris"
	
	starAddress := &star
  fmt.Printf("The address of star is %v",starAddress)
}