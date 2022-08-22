package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

type error interface {
	Error() string
}

type MyErrorType struct {
	Err   error
	debug string
}

func (e *MyErrorType) Error() string {
	return fmt.Sprintln(e.Err, e.debug)

}

var (
	temp        = errors.New("Причина остановки функции")
	errOpenFile MyErrorType
	errDev      MyErrorType
)

func OpenFile(path string) error {

	defer func() {

		if v := recover(); v != nil {
			fmt.Println("Функция OpenFile была восстановлена")
			errOpenFile.Err = temp
			errOpenFile.Err = fmt.Errorf("%w: %v", errOpenFile.Err, v)
			errOpenFile.debug = string(debug.Stack())
		}
	}()
	_, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic("OpenFile - имеет ошибку.")

	}
	return err

}

func Dev(a, b int) int {
	defer func() {

		if v := recover(); v != nil {
			fmt.Println("Функция Dev была восстановлена")
			errDev.Err = temp
			errDev.Err = fmt.Errorf("%w: %v", errDev.Err, v)
			errDev.debug = string(debug.Stack())
		}
	}()
	res := a / b
	return res
}

func fourDZ() {

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}

func main() {
	// первое и второй дз
	_ = OpenFile("data.jason")
	fmt.Println(errOpenFile)
	fmt.Println("Приложение работает дальше")

	fmt.Println(Dev(9, 0))
	fmt.Println(errDev)
	fmt.Println("Приложение работает дальше")
	// четвертое дз
	fourDZ()
	fmt.Println("Приложение работает дальше")
}
