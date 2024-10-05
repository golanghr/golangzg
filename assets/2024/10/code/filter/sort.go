package main

import (
	"sort"
)

func (v Warehouse) SortedByName() func(func(string, Grocery) bool) {
	return func(continueLoop func(string, Grocery) bool) {
		names := []string{}
		for _, v := range v.Groceries {
			names = append(names, v.Name)
		}
		sort.Strings(names)
		for _, k := range names {
			for serial, v := range v.Groceries { // example only
				if v.Name != k {
					continue
				}
				if !continueLoop(serial, v) {
					return
				}
			}
		}
	}
}
