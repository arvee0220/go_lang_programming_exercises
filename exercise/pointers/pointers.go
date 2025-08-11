//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	Name string
	TagActive bool
}

func toggleStatus(item *Item) {
	item.TagActive = !item.TagActive
}

func checkout(items *[]Item) {
	for i := range *items {
		if !(*items)[i].TagActive {
			continue
		}

		toggleStatus(&(*items)[i])
	}
}

func main() {
	items := []Item{
		{Name: "Laptop", TagActive: true},
		{Name: "Smartphone", TagActive: true},
		{Name: "Tablet", TagActive: true},
		{Name: "Headphones", TagActive: true},
	}

	fmt.Println("Initial items:", items)

	// Deactivate the tag of the first item
	toggleStatus(&items[0])
	fmt.Println("After deactivating first item:", items)

	// Call checkout to deactivate all tags
	checkout(&items)
	fmt.Println("After checkout (deactivating all tags):", items)

}
