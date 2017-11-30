// Napraw kod aby był poprawnie działającym FizzBuzzem

package main

import "fmt"

func main() {
	const n = 100
	for i := 0; i <= n; i++ {
		var output string
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output != "" && i != 0 {
			fmt.Println(output)
			continue
		}

		fmt.Println(i)
	}
}
