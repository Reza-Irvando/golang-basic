package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama  string `json:"name"`
	Kelamin string `json:"gender"`
	Umur int `json:"age"`
}

func main() {
	// Json
	// var jsonRequest = `{"name": "Alice", "gender": "female", "age": 30}`
	// var jsonData = []byte(jsonRequest)
	// var data User
	// var err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("Name: ", data.Nama)
	// fmt.Println("Gender: ", data.Kelamin)
	// fmt.Println("Age: ", data.Umur)

	// Json Interface
	fmt.Println("\nJson Interface")
	var jsonRequest2 = `{"name": "Bob", "gender": "male", "age": 25}`
	var jsonData2 = []byte(jsonRequest2)
	var data2 map[string]interface{}
	json.Unmarshal(jsonData2, &data2)
	fmt.Println("Name: ", data2["name"])
	fmt.Println("Gender: ", data2["gender"])
	fmt.Println("Age: ", data2["age"])

	// Json Array
	fmt.Println("\nJson Array")
	var jsonRequest3 = `[{"name": "Alice", "gender": "female", "age": 30},{"name": "Bob", "gender": "male", "age": 25}]`
	var jsonData3 = []byte(jsonRequest3)
	var data3 []User
	var err = json.Unmarshal(jsonData3, &data3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for index, person := range data3{
		fmt.Printf("%d. Nama: %s | Kelamin: %s | Umur: %d\n", index+1, person.Nama, person.Kelamin, person.Umur)
	}
	for i:= 0; i < len(data3); i++{
		fmt.Printf("%d. Nama: %s | Kelamin: %s | Umur: %d\n", i+1, data3[i].Nama, data3[i].Kelamin, data3[i].Umur)
	}
	fmt.Println("\nData: ")
	fmt.Println(data3)
}