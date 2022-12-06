package main

import (
	"fmt"
	"time"
)

func main() {
	var question = "We happy?"

	fmt.Println(question)

	var time = time.Now().Weekday()

	fmt.Println("Today is", time)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("tamanho:", len(b))
	fmt.Println("Esse e o array", b)
}
