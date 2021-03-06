package main

import "fmt"

func main() {
	var a complex128 = complex(6, 2)
	var b complex128 = complex(9, 2)
	var d complex64 = complex(3, 5)
	c := a + b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("a+b:", c)
	fmt.Println("d:", d)

	// Display the type
	fmt.Printf("The type of a is %T, type of b is %T and "+"the type of d is %T", a, b, d)
}
