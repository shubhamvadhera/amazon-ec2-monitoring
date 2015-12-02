package main

import "fmt"

//calFibonacci fuction calculates Fibonacci Series recursively
func calFibonacci(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return (calFibonacci(num-1) + calFibonacci(num-2))
	}
}

func main() {
	for i := 1; ; i++ {
		fmt.Println(calFibonacci(i))
	}
}
