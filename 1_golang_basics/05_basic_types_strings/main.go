package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "Sambhav"
	lastName := "Thakkar"
	
	fullName := firstName + " " + lastName
	
	fmt.Println(fullName)
	
	fmt.Println(strings.ToUpper(fullName))
	
}