package main

import "fmt"

func main() {
	sentence := "Golang The Best Language"

	fmt.Println("Soal 1 : ")
	for index, letter := range sentence {
		if index%2 == 0 {
			fmt.Println("letter :", string(letter))
		}
	}
	fmt.Println("")
	fmt.Println("Soal 2 : ")
	for _, letter := range sentence {
		letterString := string(letter)
		switch letterString {
		case "a", "e", "i", "o", "u":
			fmt.Println("letter :", string(letter))
		}
	}
}
