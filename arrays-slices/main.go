package main

import (
	"encoding/json"
	"fmt"
)
func main() {
	//arrays are copied when reassigning them to a new variable and doesnt affect the original array
	//arrays are values and do are not pointers to the first element of the array
	sample1 := [2]string{"a", "b"}
    fmt.Printf("Sample1 Before: %v\n", sample1)
    sample2 := sample1
    sample2[0] = "c"
    fmt.Printf("Sample1 After assignment: %v\n", sample1)
    fmt.Printf("Sample2: %v\n", sample2)
    test(sample1)
    fmt.Printf("Sample1 After Test Function Call: %v\n", sample1)

	//slicing of slices is not inclusive of the last item meaning it goes from start to end-1 where arr[start:end]
	fmt.Println("=====================================Slicing of slices")
	
	slicingOfSlices()
	fmt.Println("=====================================Array make")
	arrayMake()
	fmt.Println("=====================================append slice")
	appendSlice()
	fmt.Println("=====================================convert slice to json")
	convertSliceToJson()

}
func test(sample [2]string) {
    sample[0] = "d"
    fmt.Printf("Sample in Test function: %v\n", sample)
}
func slicingOfSlices(){
	numbers := [5]int{1, 2, 3, 4, 5}

    //Both start and end
    num1 := numbers[2:4]
    fmt.Println("Both start and end")
    fmt.Printf("num1=%v\n", num1)
    fmt.Printf("length=%d\n", len(num1))
    fmt.Printf("capacity=%d\n", cap(num1))

    //Only start
    num2 := numbers[2:]
    fmt.Println("\nOnly start")
    fmt.Printf("num1=%v\n", num2)
    fmt.Printf("length=%d\n", len(num2))
    fmt.Printf("capacity=%d\n", cap(num2))

    //Only end
    num3 := numbers[:3]
    fmt.Println("\nOnly end")
    fmt.Printf("num1=%v\n", num3)
    fmt.Printf("length=%d\n", len(num3))
    fmt.Printf("capacity=%d\n", cap(num3))

    //None
    num4 := numbers[:]
    fmt.Println("\nOnly end")
    fmt.Printf("num1=%v\n", num4)
    fmt.Printf("length=%d\n", len(num4))
    fmt.Printf("capacity=%d\n", cap(num4))
}

func arrayMake(){
	numbers := make([]int, 3, 5)
    fmt.Printf("numbers=%v\n", numbers)
    fmt.Printf("length=%d\n", len(numbers))
    fmt.Printf("capacity=%d\n", cap(numbers))

    //With capacity ommited
    numbers = make([]int, 3)
    fmt.Println("\nCapacity Ommited")
    fmt.Printf("numbers=%v\n", numbers)
    fmt.Printf("length=%d\n", len(numbers))
    fmt.Printf("capacity=%d\n", cap(numbers))
}
func appendSlice(){
	numbers1 := []int{1, 2}
    numbers2 := []int{3, 4}
    numbers := append(numbers1, numbers2...)
    fmt.Printf("numbers=%v\n", numbers)
    fmt.Printf("length=%d\n", len(numbers))
    fmt.Printf("capacity=%d\n", cap(numbers))
}
func convertSliceToJson(){
	a := make([]string, 2)
    a[0] = "John"
    a[1] = "Sam"
    j, err := json.Marshal(a)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    } else {
        fmt.Println(string(j))
    }
}