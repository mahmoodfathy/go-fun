package math

import "fmt"

//init functions are called first
//usually done for some first time initialization like a network call or db connect
func init() {
	fmt.Println("In add file")
}
func Add(a int, b int) int {
	return a + b
}
// func SubtractFromOtherFile(a int, b int) int {
// 	return Subtract(a, b)
// }