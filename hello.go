package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("This text brought to you by an imported package: " + quote.Go())
}
