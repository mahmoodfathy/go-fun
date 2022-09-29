package main

import "fmt"

func main() {
	
	var b *int
	fmt.Print("Default Zero Value of a pointer: ")
    fmt.Println(b)
	a := 2
	b = &a
	c := &b // pointer to pointer

	fmt.Println(**c)
	
	//Will print a address. Output will be different everytime.
	fmt.Println(b)

	fmt.Println(a)
	fmt.Println(*b)
	//declaring a pointer can be done using the new keyword
	b = new(int)
	*b = 10
	fmt.Println(*b)
	// b = b + 1 not allowed to do pointer arithmetic 
}