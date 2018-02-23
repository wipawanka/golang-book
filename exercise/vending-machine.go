package main

import (
	"fmt"
)

type NewVendingMachine struct {
	coins string
	items string
}


func (vm NewVendingMachine) InsertCoin( cointype string) string {
	return "T"
}

func (vm NewVendingMachine) GetInsertedMoney() float64 {
	
	return 10
}

func main() {
	
	vm := NewVendingMachine{coins,items}
	// Buy SD(soft drink) with exact change  
	vm.InsertCoin("T")  
	//vm.InsertCoin("F")  
	//vm.InsertCoin("TW")  
	//vm.InsertCoin("O")  
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18  
	//can := vm.SelectSD()  
	//fmt.Println(can) // SD  
	/*
	// Buy CC(canned coffee) without exact change  
	vm.InsertCoin("T")  
	vm.InsertCoin("T")  
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20  
	can = vm.SelectCC()  
	fmt.Println(can) // CC, F, TW, O  

	// Start adding change but hit coin return  
	vm.InsertCoin("T")  
	vm.InsertCoin("T")  
	vm.InsertCoin("F")  
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25  
	change := vm.CoinReturn()  
	fmt.Println(change) // T, T, F 
	*/ 
}



