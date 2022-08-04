package main

import (
	MyPackage "GO01/Golang/envconfig/MyPackage"
	"fmt"
	"os"
)

func main() {
	MyPackage.MyConfig()

	port, _ := os.LookupEnv("PORT")
	fmt.Println("Заменили значения порта на ", port)

}
