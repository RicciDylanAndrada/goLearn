package main

import "fmt"

func sums(nums ...int) {
	fmt.Println(nums, "")
}

func Functions() {
	sums(1, 2)
	sums(1, 2, 3)
	nums := []int{4, 5, 6}
	sums(nums...)
}
