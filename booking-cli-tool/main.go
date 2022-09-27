package main

import (
	"booking-cli/validator"
	"encoding/json"
	"fmt"
)

var remainingTickets int = 50
var userInfoList[] UserInfo
var userInfoListJson[] string
func main(){
for {
	greetUsers()
	getUserInfo()
	fmt.Printf("number of remaining tickets is : %v \n",remainingTickets)
	if(remainingTickets==0){
		fmt.Println("All Tickets are sold ")

		break
	}
}

}

func greetUsers(){
	fmt.Println("Welcome to the Go conference , You Can book your tikcets here")
	fmt.Println("===============================================================")
	fmt.Printf("number of remaining tickets is : %v \n",remainingTickets)
}
type UserInfo struct{
	FirstName string `json:"firstname"`
	LastName string  `json:"lastname"`
	Email string     `json:"email"`
	NumberOfTickets int  `json:"numberOfTickets"`
}
func getUserInfo(){
var userInfo UserInfo 
var firstName string
var lastName string
var email string
var numberOfTickets int
fmt.Println("Enter your first name:")
fmt.Scanln(&firstName)
userInfo.FirstName=firstName
fmt.Println("Enter your last name:")
fmt.Scanln(&lastName)
userInfo.LastName=lastName
fmt.Println("Enter your email:")
fmt.Scanln(&email)
isValidEmail := validator.ValidateEmail(email)
if(!isValidEmail){
fmt.Println("Please Enter a valid email")
	return
}
userInfo.Email=email
fmt.Println("Enter your the number of tickets you want to book:")
fmt.Scanln(&numberOfTickets)
isValidTickets := validator.ValidateTickets(remainingTickets,numberOfTickets)
if(!isValidTickets){
	fmt.Printf("number of tickets booked should be smaller than the remaining tickets remaining is %v \n",remainingTickets)
	return 
}
remainingTickets=remainingTickets-numberOfTickets

userInfo.NumberOfTickets = numberOfTickets
userInfoJson,_ :=json.Marshal(&userInfo)
userInfoListJson = append(userInfoListJson, string(userInfoJson))
// b :=json.NewEncoder(os.Stdout).Encode(&userInfoJson)
fmt.Printf("The info booked is %v \n",userInfoListJson)
}