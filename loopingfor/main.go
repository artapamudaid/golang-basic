package main

import (
	"fmt"
)

func main() {

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya Belajar Go :", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya Belajar Go :", i)
	// 	i++
	// }

	title := "Golang is the best"

	for _, letter := range title {
		fmt.Println("letter :", string(letter))
	}
}
