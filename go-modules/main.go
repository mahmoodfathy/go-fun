package main

import (
	"fmt"

	"github.com/pborman/uuid"

	"go-modules/math"
)

//order of execution

//1.Global variables in package b will be initialized. init function in source files of package b will be run
//2.Global variables in the package a will be initialized. init function in source files of package b will be run
//3.Global variables in the main package will be initialized. init function in source files of the main package will be run
//4.The main function will start executing.
func main(){
printUUID()
a:=10
b:=11
result1 := math.Subtract(a,b)
result2 :=math.Add(a,b)

fmt.Printf("result of subtraction is: %v, and result of addition is: %v ",result1,result2)
//expected output
//  "in add file"
//	"in subtract file"
//   new uuid
//result of subtraction is: -1, and result of addition is: 21 "


}

func printUUID(){
	uuid := uuid.NewRandom()
	fmt.Printf("new uuid %v \n",uuid)
}
