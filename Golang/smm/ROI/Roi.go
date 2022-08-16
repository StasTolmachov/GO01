package main

import "fmt"

func main() {

	var cpc float64            // стоимость за клик
	var clickSumBought float64 // колличесвто купленных кликов
	var ltv float64            // lifeTime Value
	var userSum float64        // колличество зарегестрированных пользователей
	var roi float64            // returne of investment
	var income float64         // доход от вложений
	var investment float64     // сумма вложений
	var registration float64   // конверсия в регистрацию

	ltv = 15.0
	clickSumBought = 3500.0
	registration = 30.0
	cpc = 4.0
	userSum = (clickSumBought / 100.0) * registration // 3500 / 100 * 30 = 1050

	income = ltv * userSum            // 15 * 1050 = 15750
	investment = cpc * clickSumBought // 4 * 3500 = 14000

	roi = Roi(income, investment) // ((15750 - 14000) / 14000) * 100 = 12.5
	fmt.Println("ROI", roi)

}

// Roi returne of investment
func Roi(inc, inv float64) float64 {
	roi := ((inc - inv) / inv) * 100
	return roi
}
