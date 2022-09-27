package main

import "fmt"
func main(){
printHelloWorld()
}

func printHelloWorld(){
hello := "Hello" // short hand
const world string = "world" // longer version
fmt.Printf("%v %v \n",hello,world)
}