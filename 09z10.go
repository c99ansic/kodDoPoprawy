package main

import "fmt"

func main() {
	for i := 3; i != 100; {
		fmt.Println("Aby wyjść podaj liczbę 100")
		fmt.Scanf("%d", &i)
	}
}
