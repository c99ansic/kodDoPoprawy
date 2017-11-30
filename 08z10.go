// Program zakłada, że każdy miesiąc ma 31 dni oraz
// poda pierwsz znak w kolejności, jeżeli w danym dniu zaczynają się dwa znaki
package main

import "fmt"

func main() {
	var (
		dzień   byte
		miesiąc byte
		wynik   string
	)
	fmt.Println("Podaj dzień urodzenia: ")
	fmt.Scanf("%d", &dzień)
	fmt.Println("Podaj miesiąc urodzenia: ")
	fmt.Scanf("%d", &miesiąc)

	if dzień > 31 || dzień < 1 {
		fmt.Println("Błędny dzień")
		return
	}

	if miesiąc > 12 || miesiąc < 1 {
		fmt.Println("Błędny miesiąc")
		return
	}

	if (miesiąc == 2 && dzień >= 16) || (miesiąc == 3 && dzień <= 11) {
		wynik = "Wodnika"
	} else if (miesiąc == 3 && dzień >= 11) || (miesiąc == 4 && dzień <= 18) {
		wynik = "Ryby"
	} else if (miesiąc == 4 && dzień >= 18) || (miesiąc == 5 && dzień <= 13) {
		wynik = "Barana"
	} else if (miesiąc == 5 && dzień >= 13) || (miesiąc == 6 && dzień <= 21) {
		wynik = "Byka"
	} else if (miesiąc == 6 && dzień >= 21) || (miesiąc == 7 && dzień <= 20) {
		wynik = "Bliźnięta"
	} else if (miesiąc == 7 && dzień >= 20) || (miesiąc == 8 && dzień <= 10) {
		// XD
		wynik = "Raka"
	} else if (miesiąc == 8 && dzień >= 10) || (miesiąc == 9 && dzień <= 16) {
		wynik = "Lwa"
	} else if (miesiąc == 9 && dzień >= 16) || (miesiąc == 10 && dzień <= 30) {
		wynik = "Pannę"
	} else if (miesiąc == 10 && dzień >= 30) || (miesiąc == 11 && dzień <= 23) {
		wynik = "Wagę"
	} else if (miesiąc == 11 && dzień >= 23) || (miesiąc == 11 && dzień <= 29) {
		wynik = "Skorpiona"
	} else if (miesiąc == 12 && dzień >= 29) || (miesiąc == 12 && dzień <= 17) {
		wynik = "Wężownika"
	} else if (miesiąc == 12 && dzień >= 17) || (miesiąc == 1 && dzień <= 20) {
		wynik = "Strzelca"
	} else if (miesiąc == 1 && dzień >= 20) || (miesiąc == 3 && dzień <= 16) {
		wynik = "Koziorożca"
	}

	fmt.Println("Masz " + wynik)
}
