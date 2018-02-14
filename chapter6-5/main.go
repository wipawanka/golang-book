package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzbuzz(number))
	}
}

func fizzbuzz(number int) (str string) {
	if number%15 == 0 {
		str = "FizzBuzz"
	} else if number%3 == 0 {
		str = "Fizz"
	} else if number%5 == 0 {
		str = "Buzz"
	} else {
		str = strconv.Itoa(number)
	}
	return str
}
