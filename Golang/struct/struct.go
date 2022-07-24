package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		rost int
	}
	var katy person
	katy.name = "Katy"
	katy.age = 30
	katy.rost = 180

	fmt.Println(katy.rost)

}

/*
package main

import "fmt"

func sayHello(hello languager) {
	fmt.Println(hello.hello())
}

type languager interface {
	hello() string
}

type ukraine struct{}

func (ukraine) hello() string {
	return "привіт"
}

type english struct{}

func (english) hello() string {
	return "hello"
}
func (english) bye() string {
	return "bye"
}

func main() {

	// sayHello("privet")
	// sayHello("hello")

	ukraniane := ukraine{}
	fmt.Println(ukraniane.hello())

	fmt.Println("------")

	englisher := english{}
	fmt.Println(englisher.hello())
	fmt.Println("------")
	fmt.Println(englisher.bye())
	fmt.Println("------")

	sayHello(ukraniane)
}
*/
