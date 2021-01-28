package main

import (
	"fmt"
	"sort"
)

func sum(array []int) int {
	result := 0
	for _, sum := range array {
		result += sum
	}
	return result
}

func avg(array []int) float64 {
	result := 0
	for _, sum := range array {
		result += sum
	}
	return float64(result) / float64(len(array))
}

func minmax(array []int) (int, int) {
	min := array[0]
	max := array[0]

	for _, v := range array {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	numbers := [32]int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3}

	chunk := len(numbers) / 3
	sliceNumbers1 := numbers[:chunk]
	sliceNumbers2 := numbers[chunk : chunk*2]
	sliceNumbers3 := numbers[chunk*2:]

	fmt.Println("Membagi Data Menjadi 3 Kelompok")
	fmt.Println(sliceNumbers1)
	fmt.Println(sliceNumbers2)
	fmt.Println(sliceNumbers3)
	fmt.Println("=======================")

	sort.Sort(sort.Reverse(sort.IntSlice(sliceNumbers1)))
	sort.Sort(sort.Reverse(sort.IntSlice(sliceNumbers2)))
	sort.Sort(sort.Reverse(sort.IntSlice(sliceNumbers3)))
	fmt.Println("Mengurutkan Data Kelompok Terbesar ke Terkecil")
	fmt.Println(sliceNumbers1)
	fmt.Println(sliceNumbers2)
	fmt.Println(sliceNumbers3)
	fmt.Println("=======================")

	fmt.Println("Menghitung Total Data Tiap Kelompok")
	fmt.Println(sum(sliceNumbers1))
	fmt.Println(sum(sliceNumbers2))
	fmt.Println(sum(sliceNumbers3))
	fmt.Println("=======================")

	fmt.Println("Menghitung Rata-rata Data Tiap Kelompok")
	fmt.Println(avg(sliceNumbers1))
	fmt.Println(avg(sliceNumbers2))
	fmt.Println(avg(sliceNumbers3))
	fmt.Println("=======================")

	fmt.Println("Mencari Nilai Terendah dan Tertinggi Tiap Kelompok")
	fmt.Println(minmax(sliceNumbers1))
	fmt.Println(minmax(sliceNumbers2))
	fmt.Println(minmax(sliceNumbers3))
}
