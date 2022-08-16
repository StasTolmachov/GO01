// шахматы
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	chess := [8][8]rune{}
	// значение верх/низ
	valueTop := [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}
	valueButton := [8]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}

	// записываем верх
	for i := range chess[0] {
		chess[0][i] = valueTop[i]
	}

	// верх пешками
	for i := range chess[1] {
		chess[1][i] = 'p'
	}

	// низ
	for i := range chess[7] {
		chess[7][i] = valueButton[i]
	}

	// низ пешками
	for i := range chess[6] {
		chess[6][i] = 'P'
	}

	chess[1][2] = '-'
	chess[3][2] = 'p'

	// записали в файл расположение фигур
	bufer2, _ := json.Marshal(chess)
	os.WriteFile("chess.json", bufer2, 0777)

	// заполнение пустой доски
	for iFirstArr, valueFirstArr := range chess {
		for iSecondArr, valueSecondArr := range valueFirstArr {
			if valueSecondArr == 0 {
				chess[iFirstArr][iSecondArr] = '-'
			}
		}
	}

	// вывод
	for print := range chess {
		fmt.Printf("%c\n", chess[print])
	}



}
