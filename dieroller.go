package main

import "fmt"
import "dieroller"

func main() {
	var max int
	for i := 0; i < 1000; i++ {
		var result = dieroller.roll()
		if result > max {
			max = result
		}
	}
	fmt.Printf("Max: %v", max)
}