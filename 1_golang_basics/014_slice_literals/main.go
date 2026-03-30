package main

import (
	"fmt"
)

func main() {
	//common collection type
	//dynamic and can grow
	//[]type{...}

	results := []string{"sambhav", "garv", "prapti"}
	fmt.Println(results, results[0], results[len(results)-1])
	//slice, first element, last element

	results[1] = "gaurav"
	fmt.Println(results)

	var nums []int
	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)
	fmt.Println(nums)
}
