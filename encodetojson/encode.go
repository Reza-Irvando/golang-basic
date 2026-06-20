package main

import ("fmt"
"encoding/json")

type Product struct {
	Name string `json:"name"`
	Amount int `json:"amount"`
}

func main(){
	// Struct to JSON
	fmt.Println("Struct to JSON!")
	var object = []Product{{"Car", 10}, {"Motor", 7}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData)
	fmt.Println("Value struct:\n", object)
	fmt.Println("Value Json:\n", jsonString)
}