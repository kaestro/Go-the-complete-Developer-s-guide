package main

import "fmt"

func main() {
	nums := []int{}
	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	for _, n := range nums {
		if n % 2 == 0 {
			fmt.Println("even", n)
		} else {
			fmt.Println("odd", n)
		}
	}
}
