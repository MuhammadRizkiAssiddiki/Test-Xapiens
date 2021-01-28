package main

import (
	"fmt"
	"strings"
)

func main() {
	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio."
	text := strings.ToUpper(lorem)

	result := ""

	//MENGHITUNG JUMLAH HURUF YANG MUNCUL
	fmt.Println("MENGHITUNG JUMLAH HURUF YANG MUNCUL")
	for i := 'A'; i <= 'Z'; i++ {
		char := string(i)
		fmt.Println("Karakter", char, "sebanyak", strings.Count(text, char), "kali")
	}

	fmt.Println("=================================================")

	//MERUBAH POSISI HURUF NAIK 5 Ex: A Menjadi F, B Menjadi G, dst.
	fmt.Println("MERUBAH POSISI HURUF NAIK 5 Ex: A Menjadi F, B Menjadi G, dst.")
	fmt.Println()
	for _, v := range text {
		v += 5
		if v > 90 { //Jika sudah mencapai huruf Z, maka akan kembali ke huruf A
			v = v - 26
		}
		if v == 37 { //Jika ada spasi maka tidak dirubah
			v = 32
		}
		if v == 49 { //Jika ada koma, maka tidak dirubah
			v = 44
		}
		if v == 51 { //Jika ada titik, maka tidak dirubah
			v = 46
		}
		result = string(v)
		fmt.Print(result)
	}

}
