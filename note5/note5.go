/*
package main

import "fmt"

func main() {
	var a uint32 = 1
	var b *uint32 = &a
	var c uint32 = *b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()

	//c++

	fmt.Println("............")
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var f *uint32 = &c
	fmt.Println(f)
}

*/
/*
package main

import "fmt"

type User struct {
	Name    string
	Surname string
	Phone   string
}

// привязываем метод к типу *User
func (u *User) SetPhone(phone string) error {
	if err := validatePhone(phone); err != nil {
		return err
	}
	u.Phone = phone
	return nil
}
func validatePhone(phone string) error {
	// логика валидации
	return nil
}
func main() {
	user := User{Name: "Иван", Surname: "Иванов"}
	if err := user.SetPhone("+79165230808"); err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v\n", user)
}
*/
/*
package main

import "fmt"

func main() {

		a := 33
		fmt.Println(a)
		fmt.Println(&a)

		b := &a
		fmt.Println(b)
		fmt.Println(*b)
		fmt.Printf("%T\n", b)

	canada := "Canada"

	var home *string

	home = &canada

	fmt.Println(home)
	fmt.Println(*home)
}
*/
/*
package main

import "fmt"
func main() {
type person struct {
	name, superpower string
	age int
}

katy := &person{
	name: "katerina",
	age: 30,
}

katy.superpower = "boom"
fmt.Println(katy)
}
*/

package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {

	rebecca := person{
		name:       "rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)


	fmt.Printf("%+v\n", rebecca)
}
