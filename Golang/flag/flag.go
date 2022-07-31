// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// )

// var (
// 	maxRows = flag.Int("max-rows", 5, "Максимальное число строк, которое нужно прочитать")
// 	colums  = flag.String("colums", "", "Разделенный запятыми список столбцов, которые нужно вывести")
// )

// func main() {
// 	flag.Parse()
// 	if flag.NArg() != 1 {
// 		log.Fatal("Неверно задано колличество аргументов")
// 	} else {

// 		fmt.Println(maxRows)
// 	}
// }
/*
package main

import (
	"fmt"
	"os"
)

func main() {
	type bd struct {
		userName string
		pass     string
		host     string
		port     string
	}
	var a bd

	a.userName = "stasvv"
	a.pass = "password"
	a.host = "www"
	a.port = "8080"
	fmt.Println(a)

	file, err := os.Create("DB.txt")
	if err != nil {
		fmt.Println("файл не создан", err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(a.userName)

}
*/

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	text := "Hello Gold!"
// 	file, err := os.Create("hello.txt")

// 	if err != nil {
// 		fmt.Println("Unable to create file:", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
// 	file.WriteString(text)

// 	fmt.Println("Done.")
// }

package main

import (
	"flag"
	"fmt"
)

type baseData struct {
	userName string
	pass     string
	host     string
	port     int
}

func main() {
	connect := baseData{}
	flag.StringVar(&connect.userName, "userName", "admin", "имя пользователя")
	flag.StringVar(&connect.pass, "pass", "admin", "пароль")
	flag.StringVar(&connect.host, "host", "0.0.0.0", "хост")
	flag.IntVar(&connect.port, "port", 8080, "порт")

	flag.Parse()

	flagFunc(connect.userName, connect.pass, connect.host, connect.port)
}
func flagFunc(userName, pass, host string, port int) {
	fmt.Println(userName, pass, host, port)
}
