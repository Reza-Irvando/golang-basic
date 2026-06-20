package main

import (
	"fmt"
	"runtime"
)
func main(){
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello, %s", who)
		messages <-  data
	}
	go sayHelloTo("Dono")
	go sayHelloTo("Kasino")
	go sayHelloTo("Indro")
	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)
}