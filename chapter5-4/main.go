package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i int
	rand.Seed(time.Now().UTC().UnixNano())
	i = randInt(1, 30)
	/*
	fmt.Printf("Random number is ", i)
	*/
	for number := 1; number <= 5; number++ {
		fmt.Print("Enter magic number: ")
		var input int
		fmt.Scanln(&input)
		if input < i {
			fmt.Printf("Less Than \n")
		} else if input > i {
			fmt.Printf("Too Much \n")
		} else {
			fmt.Printf("Equal \n")
		}
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
