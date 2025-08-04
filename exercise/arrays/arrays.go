//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price float64
}

var shoppingList = [4]Product{
	{ name: "Apples", price: 1.20 },
	{ name: "Bread", price: 2.50 },
	{ name: "Milk", price: 1.80 },
}

func printShoppingList(list [4]Product) {

	if list[len(list) - 1].name == "" && list[len(list) - 1].price == 0 {
		fmt.Println("Last item on the list:", list[len(list) - 2].name, " - $", list[len(list) - 2].price)
	} else {
		fmt.Println("Last item on the list:", list[ len(list) - 1].name, " - $", list[ len(list) - 1].price)
	}
	
	fmt.Println("Total number of items:", len(list))

	totalCost := 0.0

	for _, item := range list {
		totalCost += item.price
	}

	fmt.Println("Total cost of items:", totalCost)
}

func main() {

	printShoppingList(shoppingList)

	shoppingList[3] = Product{ name: "Eggs", price: 2.00 }

	printShoppingList(shoppingList)

}

