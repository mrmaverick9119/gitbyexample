package main

import (
	"example/greeting"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(greeting.Hello("napier"))
}
