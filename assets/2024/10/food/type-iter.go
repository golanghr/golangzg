package main

import (
	"sort"
)

func (v Warehouse) SortedByName() func(func(int, Grocery) bool) {
	return func(yield func(int, Grocery) bool) {
		index := 0
		keys := []string{}

		for _, v := range v.Groceries {
			keys = append(keys, v.Name)
		}
		sort.Strings(keys)
		for _, k := range keys {
			for _, v := range v.Groceries {
				if v.Name != k {
					continue
				}
				if !yield(index, v) {
					return
				}
				index++
			}
		}
	}
}

func (v Warehouse) Filter(foodType string) func(func(Grocery) bool) {
	return func(yield func(Grocery) bool) {
		for _, v := range v.Groceries {
			if v.Type != foodType {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}
