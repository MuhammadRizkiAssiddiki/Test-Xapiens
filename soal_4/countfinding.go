package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := randInt(1, 100)
	number := randInt(1, 100)
	count := 1

	fmt.Println("Angka yang dicari :", guess)
	fmt.Println()

	fmt.Println("Percobaan ke - 1 :", number)

	for number != guess {
		number = randInt(1, 100)
		count++
		fmt.Println("Percobaan ke -", count, ":", number)
	}

	fmt.Println("Jumlah Percobaan:", count)
	fmt.Println("Suatu percobaan yang mantap (y)")
}
