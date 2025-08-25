package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")

	if !found {
		fmt.Println("No, we don't need cereal")
	} else {
		fmt.Println("Yes, we need", cereal, "cereal")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total items in the shopping list:", totalItems)
}
