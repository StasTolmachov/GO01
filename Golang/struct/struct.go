package main

func main() {
	// type person struct {
	// 	name string
	// 	age  int
	// 	rost int
	// }
	// var katy person
	// katy.name = "Katy"
	// katy.age = 30
	// katy.rost = 180

	// fmt.Println(katy.rost)

	type shape struct {
		a int
		b int
		// c = int
	}

shape.a = 4
shape.b = 5


kvadrat := shape



	func (s shape) square() {
		return int(kvadrat.a * kvadrat.b)
	} 

	sq := kvadrat.square()

	fmt.Println(sq)


	

}
