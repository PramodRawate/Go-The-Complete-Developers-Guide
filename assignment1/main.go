package main

import "fmt"

func main() {
	numbers1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("**********Solution-1***********")
	for _, num := range numbers1 {
		if num%2 == 0 {
			println(num, " is even")
		} else {
			println(num, " is odd")
		}
	}

	numbers2 := [11]int{}
	fmt.Println("**********Solution-2***********")
	for i := 0; i <= 10; i++ {
		numbers2[i] = i
		if numbers2[i]%2 == 0 {
			println(numbers2[i], " is even")
		} else {
			println(numbers2[i], " is odd")
		}
	}

	fmt.Println("**********Solution-3***********")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			println(i, " is even")
		} else {
			println(i, " is odd")
		}
	}

}
