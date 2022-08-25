package main

import "fmt"

type uah float64
type usd float64
type zloty float64

// uah
func (us usd) uah() uah {
	return uah(us * 35.5)
}
func (zl zloty) uah() uah {
	return uah(zl * 6.28)
}

//usd
func (ua uah) usd() usd {
	return usd(ua / 35.5)
}
func (zl zloty) usd() usd {
	return usd(zl * 0.21)
}

// zloty
func (us usd) zloty() zloty {
	return zloty(us / 0.21)
}

func (ua uah) zloty() zloty {
	return zloty(ua / 6.28)
}

func main() {
	var zl zloty
	//var usd usd
	var uah uah
	var valutaIN string
	var valutaOUT string
	var result float64
	fmt.Println("valuta in: zl uah usd")
	fmt.Scan(&valutaIN)
	fmt.Println("valuta out: zl uah usd")
	fmt.Scan(&valutaOUT)

	if valutaIN == "uah" && valutaOUT == "usd" {
		fmt.Println("summa")
		fmt.Scan(&uah)
		result = float64(uah.usd())
		fmt.Println("first")
		fmt.Println(uah, "гривен в долларах будет ", result)

	}
	if valutaIN == "zl" && valutaOUT == "usd" {
		result = float64(zl.usd())
		fmt.Println("second")

	}
	if valutaIN == "uah" && valutaOUT == "zl" {
		result = float64(uah.zloty())
	}

}
