package main

import "fmt"

type vehichle interface {
	engine()
	move()
}

type plane struct {
	wings string
}
type car struct {
	wheels string
}
type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l *lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l *lion) walk() {
	fmt.Println("Lion walk")
}
type human interface { //must implement bothe animal methods and speak method
	animal
	speak()
}

type employee struct {
	name string
}

func (e employee) breathe() {
	fmt.Println("Employee breathes")
}

func (e employee) walk() {
	fmt.Println("Employee walk")
}

func (e employee) speak() {
	fmt.Println("Employee speaks")
}

func main() {
var firstTestVehichle vehichle
var secondTestVehichle vehichle
firstTestVehichle = plane{wings: "wings"}
secondTestVehichle = car{wheels: "wheels"}
firstTestVehichle.engine()
secondTestVehichle.engine()
firstTestVehichle.move()
secondTestVehichle.move()

callEngine(firstTestVehichle)
callEngine(secondTestVehichle)
callMove(firstTestVehichle)
callMove(secondTestVehichle)
var a animal
	// a = lion{age: 10} ---> this is throwing an error because it has a pointer reciever function
	// a.breathe()
	// a.walk()

	a = &lion{age: 5}
	a.breathe()
	a.walk()
	
	var h human

	h = employee{name: "John"} //throws error if speak method not implemented 
	h.breathe()
	h.walk()
	h.speak()
	a= dog{age: 10}
	printTypeSwitch(a) 
	printWithTypeAssertion(a)
	a= &lion{age: 10}
	printTypeSwitch(a)
	// printWithTypeAssertion(a) // if we passed a lion it will panic 

	test("thisisstring")
    test("10")
    test(true)
}
type dog struct {
	age int
}
func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}
func printWithTypeAssertion(a animal){
	l := a.(dog) //type assertion that this will be a dog
	// val, ok := i.() can be used like this as well
	fmt.Printf("Age: %d\n", l.age)
}

func printTypeSwitch(a animal) {
	
	switch v := a.(type) {
	case *lion:
		fmt.Println("Type: lion")
	case dog:
		fmt.Println("Type: dog")
	default:
		fmt.Printf("Unknown Type %T", v)
	}


}
func (c car) engine(){
	fmt.Println("A car has an engine")
}

func (p plane) engine(){
	fmt.Println("A plane has an engine")
}

func (p plane) move(){
	fmt.Printf("A plane moves by flying with %v \n",p.wings)
}
func (c car) move(){
	fmt.Printf("A car moves on its four %v \n",c.wheels)

}

func callEngine(v vehichle){
	v.engine()
}
func callMove(v vehichle){
	v.move()
}
//empty interface can be passed any type
func test(a interface{}) {
    fmt.Printf("(%v, %T)\n", a, a)
}