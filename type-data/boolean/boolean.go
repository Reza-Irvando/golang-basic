package main

import "fmt"

func main() {
	var isRaining bool = true
	var isSunny bool = false

	isLoggedIn := true
	hasPermission := false

	// Print the values
	fmt.Println("Is it raining?", isRaining)
	fmt.Println("Is it sunny?", isSunny)
	fmt.Println("Is the user logged in?", isLoggedIn)
	fmt.Println("Does the user have permission?", hasPermission)
}