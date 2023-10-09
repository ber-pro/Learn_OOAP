package main

import (
	"OO_analyze/guitar"
	"fmt"
)

func main() {
	fmt.Println("Start initialize inventory")
	inventory := guitar.NewInventory()

	guitar1 := guitar.NewGuitar("123", "Fender", "Stratocaster", "Electric", "Maple", "Alder", 999.99)
	guitar2 := guitar.NewGuitar("456", "Gibson", "Les Paul", "Electric", "Mahogany", "Maple", 1499.99)
	guitar3 := guitar.NewGuitar("789", "Martin", "D-28", "Acoustic", "Rosewood", "Spruce", 1999.99)
	guitar4 := guitar.NewGuitar("987", "fender", "Stracastor", "electric", "Alder", "Alder", 0)

	fmt.Println("Add guitars in inventory")
	inventory.AddGuitar(guitar1)
	inventory.AddGuitar(guitar2)
	inventory.AddGuitar(guitar3)
	inventory.AddGuitar(guitar4)

	fmt.Println("Get getGuitar")

	if getGuitar, ok := inventory.GetGuitar("987"); ok {
		fmt.Println("Find Guitar with serial number: ", getGuitar.GetSerialNumber())
	} else {
		fmt.Println("Will not find a getGuitar with this serial number")
	}

}
