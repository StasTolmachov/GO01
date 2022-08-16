package main

import "fmt"

func main() {

	var cpc float64            // стоимость за клик
	var clickSumBought float64 // колличесвто купленных кликов
	// var ltv float64            // lifeTime Value
	var userSum float64 // колличество зарегестрированных пользователей
	// var roi float64            // returne of investment
	// var income float64         // доход от вложений
	// var investment float64     // сумма вложений
	var registration float64 // конверсия в регистрацию
	var cpa float64          // Cost Per Action
	var coastSum float64     // стоимость всей рекламы

	// ltv = 15.0
	clickSumBought = 3200.0
	registration = 8.5
	cpc = 0.75

	coastSum = clickSumBought * cpc
	fmt.Println("coastSum", coastSum)

	userSum = (clickSumBought / 100.0) * registration
	fmt.Println("userSum", userSum)
	cpa = coastSum / userSum
	fmt.Println("CPA", cpa)
	// магазин
	var a float32
	a = 32.0 - (32.0 / 100 * 25) // 24

	var result float32
	result = a - (a / 100 * 60)

	fmt.Println("result", result)

	// микропроцессор

}
