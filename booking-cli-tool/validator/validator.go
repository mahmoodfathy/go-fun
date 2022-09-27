package validator

import (
	"strings"
)


func ValidateEmail(email string) bool{
	return strings.Contains(email,"@")
}
func ValidateTickets(remainingTickets int,tickets int)bool{
	return remainingTickets>=tickets 
}