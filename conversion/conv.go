package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Int to float
	var a int = 10
	var b float64 = float64(a)
	fmt.Println("Nilai a: ", a)
	fmt.Println("Nilai b: ", b)
	
	// Int to string & String to int
	var text string = strconv.Itoa(a)
	var c string = "100"
	number, err := strconv.Atoi(c)
	fmt.Println("Nilai text: ", text)
	fmt.Println("Nilai c: ", c)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Nilai number: ", number)
	}

	// bool to string & String to bool
	var isTrue bool = true
	var isFalse bool = false
	var trueText string = strconv.FormatBool(isTrue)
	var falseText string = strconv.FormatBool(isFalse)
	fmt.Println("Nilai isTrue: ", isTrue)
	fmt.Println("Nilai isFalse: ", isFalse)
	fmt.Println("Nilai trueText: ", trueText)
	fmt.Println("Nilai falseText: ", falseText)
	var boolText string = "true"
	boolValue, err := strconv.ParseBool(boolText)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Nilai boolValue: ", boolValue)
	}
}