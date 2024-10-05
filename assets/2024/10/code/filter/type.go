package main

type Grocery struct {
	Name string
	Type string
}

type Warehouse struct {
	Groceries map[string]Grocery
	// ...
}
