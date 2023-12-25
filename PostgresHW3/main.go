package main

import (
	"fmt"
	"marketbasa/funcforcat"
	"marketbasa/funcforprod"
)

func main() {
	num := 0
	fmt.Println(`
	Welcome
	Select a table:
	1 - access to the category table
	2 - access to the product table
	`)
	fmt.Scan(&num)
	switch num {
	case 1:
       funcforcat.CategoryFuncs()
	case 2:
		funcforprod.ProductFuncs()
	default:
		fmt.Printf("Not command found")
	}
}
