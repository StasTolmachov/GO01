package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

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

	fmt.Println("-----------")

	slice := []string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}

	fmt.Println(slice)

	fmt.Println("-----------")

	hyperspace(slice)
	fmt.Println(strings.Join(slice, ""))
	fmt.Println("-----------")

	sort.StringSlice(slice).Sort()
	fmt.Println(slice)

}

func hyperspace(words []string) {
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
	}
}

type StringSlice []string

func (p StringSlice) Sort() {}
