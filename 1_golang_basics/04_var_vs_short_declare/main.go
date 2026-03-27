package main

import "fmt"

func main(){
	var city string
	city = "London"
	
	var name = "Sambhav" //inferred to string
	
	// := (this is a short declaration method)
	subscribers := 500000
	
	likes, comments := "100", 30
	
	fmt.Println(city, name, subscribers, likes, comments)
	
}