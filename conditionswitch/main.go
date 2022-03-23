package main

import "fmt"

func main() {
	nilai := "F"
	var komentarNetizen string

	switch nilai {
	case "A":
		komentarNetizen = "Amazing"
	case "B":
		komentarNetizen = "B Aja"
	case "C":
		komentarNetizen = "Cuma Segitu?"
	case "D":
		komentarNetizen = "Payah"
	case "E":
		komentarNetizen = "Gak Guna"
	case "F":
		komentarNetizen = "Sukurin"
	}

	fmt.Println(komentarNetizen)
}
