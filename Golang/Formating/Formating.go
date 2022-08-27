package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	CreatePhoneNumber := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// returns "(123) 456-7890"

	var PhoneNumberStringArr [10]string

	for i := range CreatePhoneNumber {
		PhoneNumberStringArr[i] = strconv.Itoa(int(CreatePhoneNumber[i]))
	}

	PhoneNumberString := "(" + strings.Join(PhoneNumberStringArr[:3], "") + ")" + " " + strings.Join(PhoneNumberStringArr[3:6], "") + "-" + strings.Join(PhoneNumberStringArr[6:], "")

	fmt.Println(PhoneNumberString)

	

}
