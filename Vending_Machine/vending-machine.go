package main

import (
	"fmt"
)

type VendingMachine struct {
	
}

func (v VendingMachine) InsertMoney() int {
	return 0
}

func main() {
	vm := VendingMachine{}
	fmt.Println("Insert Money:",vm.InsertMoney())
	// Insert Money :0
}



