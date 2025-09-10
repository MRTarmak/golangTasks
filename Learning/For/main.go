package main

import "fmt"

func main() {
	var x uint

	fmt.Scan(&x)

	for i := x; i > 0; i-- {
		fmt.Println(i)
	}
}
