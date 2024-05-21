package main

import (
	algo "dsalgo/maps"
	"fmt"
)

func main() {
	fmt.Println("welcome to ds algorithm")

	character := "ajay kumar "
	count := algo.CountCharacterOccurance(character)

	fmt.Println(count)

}
