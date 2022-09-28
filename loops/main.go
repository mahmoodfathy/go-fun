package main

import (
	"flag"
	"fmt"
)
func main() {
	fmt.Println("-l --option 'normal' ,'map' , 'infinite', 'range' ")
	 loopType := flag.String("l","range","type of loop")
	 flag.Parse()
	fmt.Printf("loop type is %v \n", *loopType)
	switch *loopType{
	// looping through arrays or slices in range:
	case "range":
		arr :=[]string{"test1","test2","test3","test4"}
		loopRange(arr)

	// looping though maps:
	case "map":
		newMap := map[string]string{"name":"bro","job":"coolasfuck","function":"letsgooooo"}
		loopMap(newMap)
	//normal loops through the array or slice:
	case "normal":
		arr :=[]string{"test@test.com","test2@test.com"}
		loopNormal(arr)
	//while loops:
	case "while":
		arr :=[]string{"test@test.com","test2@test.com"}
		whileLoop(arr)
	//infinte loops:
	case "infinite":
		loopInfinite()
	default:
		break
	}






}
func loopRange(arr []string){
for i,item :=range arr{
	fmt.Printf("index is : %v and item is %v \n",i,item)
	}
}
func loopMap(newMap map[string]string){ //doesnt guarantee iteration over map in order, WOOW , need to sort manually
	for key,value :=range newMap{
		fmt.Printf("key is : %v and value is %v \n",key,value)
	}	
}
func loopNormal(arr []string){
	for i:=0;i<len(arr);i++{
		fmt.Printf("index is : %v and item is %v \n",i,arr[i])
	}
}
func whileLoop(arr []string){
	i := 0
	for i <2{
		fmt.Printf("index is : %v and item is %v \n",i,arr[i])
		i++;
	}
}
func loopInfinite(){
	var scanned int
	for{
		fmt.Println("Enter  to exit from loop")
		fmt.Scanln(&scanned)
		
		if scanned==0 {
			break
		}
	}
}