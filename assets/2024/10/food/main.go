package main

import "fmt"

func main() {
	shoppingList := Warehouse{
		Groceries: map[string]Grocery{
			"serial_number_1": {
				Name: "banana",
				Type: "fruit",
			},
			"serial_number_2": {
				Name: "apple",
				Type: "fruit",
			},
			"serial_number_3": {
				Name: "carrot",
				Type: "vegetable",
			},
		},
	}
	for i, g := range shoppingList.SortedByName() {
		fmt.Println(i, g)
	}
	fmt.Println("Filter fruit")
	for g := range shoppingList.Filter("fruit") {
		fmt.Println(g)
	}
}
