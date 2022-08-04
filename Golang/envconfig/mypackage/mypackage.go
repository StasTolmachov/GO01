package MyPackage

import (
	"fmt"
	"os"
)

func MyConfig() {
	fmt.Println("Проверка работы моего пакета")
	os.Setenv("PORT", "0987")
}
