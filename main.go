package main

import "fmt"

func main() {
	num := 8
	fmt.Println(numberOfSteps(num))
}

func numberOfSteps(num int) int {

	count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
			count++
		} else {
			num = num - 1
			count++
		}
	}
	if num > 1 {
		numberOfSteps(num)
	}
	return count
}
