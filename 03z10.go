package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Ilość znaków napisu poniżej to:")
	msg := "hejka programistyczne świry"
	// Ilość run -> znaków
	fmt.Println(utf8.RuneCountInString(msg))
	// Ilość bajrtów
	fmt.Println(len(msg))
}