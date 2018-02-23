package main

import (
	"fmt"
)

func SubtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func main() {
	/*
	for i := 0; i < 10; i++ {
		fmt.Println (DisplayNumber(i))
	}
	*/


	
	fmt.Println(weatherCelsius(25,"Bangkok few clound"))
	fmt.Println(weatherCelsius(34,"Tak sunny"))
	fmt.Println(weatherCelsius(17,"Phuket rainy"))
	fmt.Println(weatherCelsius(9,"Chiangmai cold"))
	
}



func weatherCelsius(temper int, desc string) (str string) {
	
	if temper==25 {
		/*
		for i := 0; i < 4; i++ {
			if i== 0 {
				fmt.Println (" _ _")
			}
		*/
		fmt.Printf ("|_||_| \n") //6 width + o + C
		fmt.Printf (" _  _ \n")		// line 1
		fmt.Printf (" _||_ o\n")  // Line 2
		fmt.Printf ("|_  _| C \n") // Line 3
	}

	if temper == 34 {
		fmt.Printf (" _     \n")
		fmt.Printf (" _||_| o\n")
		fmt.Printf (" _|  |  C \n")
	}
	
	if temper == 17 {
		fmt.Printf ("    _ \n")
		fmt.Printf ("  |  | o\n")
		fmt.Printf ("  |  |  C \n")
	}

	if temper == 9 {
		fmt.Printf (" _ \n") // 1 digit = 3 + o + C 
		fmt.Printf ("|_| o\n")
		fmt.Printf (" _|  C \n")
	}
	str = DisplayDesc(desc)
	return desc
}

func DisplayDesc(desc string) (Strdesc string) {
	return desc
}

func DisplayNumber (number int) (strnumber string) {
	if number == 0 {
		var num1line1 string
		num1line1 = " _ \n"
		fmt.Printf (num1line1)
		fmt.Printf ("| |\n")
		fmt.Printf ("|_|\n")
	}
	if number == 1 {
		fmt.Printf ("   \n")
		fmt.Printf ("  |\n")
		fmt.Printf ("  |\n")
	}
	if number == 2 {
		fmt.Printf (" _ \n")
		fmt.Printf (" _|\n")
		fmt.Printf ("|_ \n")
	}
	if number == 3 {
		fmt.Printf (" _ \n")
		fmt.Printf (" _|\n")
		fmt.Printf (" _|\n")
	}
	if number == 4 {
		fmt.Printf ("   \n")
		fmt.Printf ("|_|\n")
		fmt.Printf ("  |\n")
	}
	if number == 5 {
		fmt.Printf (" _ \n")
		fmt.Printf ("|_ \n")
		fmt.Printf (" _|\n")
	}
	if number == 6 {
		fmt.Printf (" _ \n")
		fmt.Printf ("|_ \n")
		fmt.Printf ("|_|\n")
	}
	if number == 7 {
		fmt.Printf (" _ \n")
		fmt.Printf ("  |\n")
		fmt.Printf ("  |\n")
	}
	if number == 8 {
		fmt.Printf (" _ \n")
		fmt.Printf ("|_|\n")
		fmt.Printf ("|_|\n")
	}
	if number == 9 {
		fmt.Printf (" _ \n")
		fmt.Printf ("|_|\n")
		fmt.Printf (" _|\n")
	}
	return strnumber
}





