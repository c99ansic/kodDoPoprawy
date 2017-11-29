package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "parzysta")
		} else {
			fmt.Println(i, "nieparzysta")
		}
	}
}
