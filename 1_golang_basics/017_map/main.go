package main

import (
	"fmt"
)

func main() {
	//map[keyType]valueType
	ages := map[string]int{
		"sambhav": 21,
		"john":    35,
	}
	fmt.Println(ages["sambhav"], len(ages))

	//using make
	//make(map[K]V)
	//
	var scores map[string]int //nill map

	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)
	scores["math"] = 90

	fmt.Println(scores)
	
	users := map[string]string{
		"u1" : "sambhav",
		"u2" : "john",
		"u3" : "roma",
 	}
  
  	fmt.Println(users)
   
   	delete(users, "u2")
    delete(users, "u4") // this will not do anything, no error
    fmt.Println(users)
}
