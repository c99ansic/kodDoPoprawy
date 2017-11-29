package main

import (
	"fmt"
	"strconv"
)

// Uzupełnij program aby wypisał poniższą linijkę
// 5 metrów to 16.404199475065617 stóp
// Wartość ma zostać obliczona wykorzystując 
// zmienne 1 stopa to 0.3048 metra


func main() {
	meters := 5

	fmt.Println(
		strconv.Itoa(meters) +
		` metrów to ` +  
		getFeets(meters) +
		` stóp`)
}

func getFeets(meters int) string {
	const feetToMeters = 0.3048
	return strconv.FormatFloat(float64(meters) / feetToMeters, 'f', 6, 64)
}

