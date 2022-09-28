package main

import (
	"fmt"
	"maps-structs/convertions"
	"reflect"
	"time"
)

//Maps can only hold one type of keys and values
//Structs can hold multiple types
func main(){

//one way to initialize the map
//map can't be nil have to initalize it to empty
var newMap map[string]int = map[string]int{}
newMap["John Doe"] = 5000
newMap["Alice"] = 3000

fmt.Printf("Salaries map is : %v \n",newMap)

//another way to initialize maps
salariesMap :=make(map[string]int)

salariesMap["Ahmed"] =1000
salariesMap["test"] =2000
fmt.Printf("Salaries map is : %v \n",salariesMap)


fmt.Println("======================================= Map converted to json")
convertion:=convertions.MapToJson(salariesMap)
fmt.Printf("conversion is : %v \n",convertion)
reverseConvertion:=convertions.JsonToMap([]byte(convertion))
fmt.Printf("reverse conversion is : %v \n",reverseConvertion)

//create a struct 
//export members to be able to convert to json
type Employee struct{
	Name string `json:"name"` //annotate to change the key name otherwise it will default to the capitlization
	Department string `json:"department"`
	Salary int  `json:"salary"`
	Time time.Time `json:"createdAt"`
}
type student struct{
	Name string `json:"name"` //annotate to change the key name otherwise it will default to the capitlization
	Age string `json:"age"`
	Universities []string
}
type Admin struct{
	Name string `json:"name"` //annotate to change the key name otherwise it will default to the capitlization
	Department string `json:"department"`
	Salary int  `json:"salary"`
	Time time.Time `json:"createdAt"`
	Admin bool `json:"admin"`
	bool  `json:"boolean"`// anonymous field will be named as the type
}

newEmployee :=&Employee{Name:"Mahmood",Department: "IT", Salary: 5000000, Time: time.Now() } //pointer to struct
newAdmin :=&Admin{Name:"Mahmood",Department: "IT", Salary: 5000000, Time: time.Now(),Admin: true,bool: false }

fmt.Printf("Employee is : %+v \n",*newEmployee) //pretty print
fmt.Printf("admin is : %+v \n",*newAdmin)//derefrence it
employeeConversion := convertions.StructToJson(newEmployee)
adminConversion := convertions.StructToJson(newAdmin) 
fmt.Printf("json Employee is : %v \n",employeeConversion)
fmt.Printf("json admin is : %v \n",adminConversion)

//struct equality

//structs can equal each other if they dont' contain types that can't be compared
employee1 :=Employee{Name:"Mahmood",Department: "IT", Salary: 5000000, Time: time.Now() }
employee2 :=Employee{Name:"Mahmood",Department: "IT", Salary: 5000000, Time: time.Now() }
isEqual := employee1==employee2
isEqual2:=reflect.DeepEqual(employee1,employee2)
fmt.Println(isEqual,isEqual2)
student1 :=student{Name:"Mahmood",Age: "10",Universities: []string{"test"} }
student2 :=student{Name:"Mahmood",Age: "10",Universities: []string{"test"} }

// areStudentsEqual :=student1==student2 // cannot be compared throws compile error

areStudentsEqualDeep := reflect.DeepEqual(student1,student2); //deep equality works and returns true
fmt.Println(areStudentsEqualDeep)

//structs are value types and assigning to a new variable will not affect the original struct
}
