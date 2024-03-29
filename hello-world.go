package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	a := "Hello World"
	fmt.Println(a)

	for i := range 3 {
		println("1 + "+strconv.Itoa(i)+" =", 1+i)
	}

	nums := []int{rand.Intn(99), rand.Intn(99), rand.Intn(99)}
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	for i := range nums {
		if nums[i]%2 == 0 {
			fmt.Println(nums[i], "is even")
		} else {
			fmt.Println(nums[i], "is odd")
		}
	}

	if num := rand.Intn(20); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
