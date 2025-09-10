package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x > 10 {
		fmt.Println("Number is higher than 10")
	} else if x == 10 {
		fmt.Println("Number equals 10")
	} else {
		fmt.Println("Number is lower than 10")
	}
}
