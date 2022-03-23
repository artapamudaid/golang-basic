package main

import "fmt"

func main() {
	age := 8
	var notice string

	if age > 10 {
		notice = "Boleh bermain HP"
	} else if age > 6 && age <= 10 {
		notice = "Boleh Main HP sebentar"
	} else {
		notice = "Belum Boleh main HP"
	}

	fmt.Println(notice)
}
