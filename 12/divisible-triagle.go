package main

import "fmt"

func main() {
	triangle := 0
	i := 1
	num := 0
	for num < 500 {
		triangle += i
		i++
		n := numDevisors(triangle)
		if n > num {
			num = n
			fmt.Printf("Triangle: %d, numDevisors: %d\n", triangle, num)
		}
	}
}

func numDevisors(n int) int {
	num := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			num++
		}
	}
	return num
}
