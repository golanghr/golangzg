package main

func (v Warehouse) Filter(foodType string) func(func(Grocery) bool) {
	return func(continueLoop func(Grocery) bool) {
		for _, v := range v.Groceries {
			if v.Type != foodType {
				continue
			}
			if !continueLoop(v) {
				return
			}
		}
	}
}
