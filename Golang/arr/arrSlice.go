package main

import (
	"fmt"
)

func main() {
	fmt.Println("work")
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(planets, terrestrial, gasGiants, iceGiants)

}
