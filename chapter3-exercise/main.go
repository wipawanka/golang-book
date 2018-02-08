package main

import "fmt"

func main() {
	fmt.Println("=====Value Expression=====")
	bolleantext :=  (true && false) || (false && true) || !(false && false) 
	fmt.Printf("bollean =%v\n",bolleantext)	
}
