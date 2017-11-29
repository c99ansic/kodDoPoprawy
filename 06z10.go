package main

import "fmt"

func main() {
	var (
		i     byte = 1
		limit byte
	)

	fmt.Println("Do ilu odliczyć?")
	fmt.Scanf("%d", &limit)
	for i <= limit {
		fmt.Println(i)
		i++
	}
}
