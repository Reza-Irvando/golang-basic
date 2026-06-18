package main

import "fmt"

func main() {
	// Create a map with string keys and int values
	map1 := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println("Map1:", map1)

	// Accessing a value
	fmt.Println("Value of apple:", map1["apple"])
	
	// create a map without int keys and string values
	map2 := make(map[string]int)

	// Adding a new key-value pair
	map2["Reza"] = 100
	map2["Hytam"] = 90
	fmt.Println("Map2 after adding:", map2)
	
	// Updating a value
	map2["Reza"] = 95

	// Deleting a key-value pair
	delete(map2, "Reza")
	fmt.Println("Map2 after update:", map2)
	fmt.Println("Map Length:", len(map2))

	// Checking if a key exists
	value, exists := map2["Hytam"]
	fmt.Println("Value of Hytam:", value)
	fmt.Println("Does Hytam exist?", exists)
}