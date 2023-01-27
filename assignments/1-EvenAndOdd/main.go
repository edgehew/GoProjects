
package main

import "fmt"

func main() {
	nums := make([]int, 11)
	for i, _ := range nums {
		nums[i] = i
	}

	for _, val := range nums {
		property := "even"
		if val % 2 != 0  {
			property = "odd"
		}

		fmt.Println(val, " is ", property) 
	}
}