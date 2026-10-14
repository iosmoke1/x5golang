package main

import "fmt"

func FizzBuzz(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	FizzBuzz(n)
}