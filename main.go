package main

import (
	two_sum "blind75/1_two_sum"
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
    nums := []int {2,7,11,15}
    target := 9 
    result := two_sum.TwoSum(nums,target)

    fmt.Printf(result)
