package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
	address 
}
type address struct {
	city    string
	country string
}

func main() {
	

    emp := employee{name: "Sam", age: 31, salary: 2000}
    emp.setNewName("John")
    //not affected by method since it created a copy of it and didnt modify the original struct
	fmt.Printf("Name: %s\n", emp.name) 
	//fix with pointer reciever
	emp.setNewNameWithPointerReciever("mike")
	fmt.Printf("Name after update: %s\n", emp.name)

	//Methods on Anonymous nested struct fields
	fmt.Println("Methods on Anonymous nested struct fields")
	address := address{city: "London", country: "UK"}
	emp2 := employee{name: "Sam", age: 31, salary: 2000, address: address}
	// two ways to access the address details , either from employee or address direct
	emp2.cityDetails() 

	emp2.address.cityDetails()

	//Method Chaining
	fmt.Println("Methods chaining by returning the struct")
	emp.printAge().printName().printSalary()
}

func (a address) cityDetails() {
	fmt.Printf("City: %s\n", a.city)
	fmt.Printf("Country: %s\n", a.country)
}
func (e employee) details() {
	fmt.Printf("Name: %s\n", e.name)
	fmt.Printf("Age: %d\n", e.age)
}

func (e employee) getSalary() int {
	return e.salary
}
func (e employee) setNewName(newName string) {
	e.name = newName
}
func (e *employee) setNewNameWithPointerReciever(newName string) {
	e.name = newName
}


func (e employee) printName() employee {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("Salary: %d\n", e.salary)
}