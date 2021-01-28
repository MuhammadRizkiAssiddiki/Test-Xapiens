package main

import "fmt"

func main() {
	for i := 1; i <= 7; i++ {
		for j := 7; j >= i; j-- {
			if j == 7 {
				fmt.Print(i)
			} else {
				fmt.Print(0)
			}

		}
		fmt.Println()
	}
}
