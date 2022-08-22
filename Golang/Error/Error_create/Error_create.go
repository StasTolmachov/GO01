// package main

// import "fmt"

// type HttpError struct {
// 	code    int
// 	message string
// }

// func (e HttpError) Error() string {
// 	return fmt.Sprintln(e.code, e.message)
// }

// func FuncHttpError(code int, message string) HttpError {
// 	return HttpError{code: code, message: message}
// }

// var (
// 	ErrNotFound  = FuncHttpError(404, "not found")
// 	ErrForbidden = FuncHttpError(403, "forbidden")
// )

// func GetPage(path string) error {
// 	switch path {
// 	case "/passwd":
// 		return ErrForbidden
// 	case "/user/1":
// 		return nil
// 	default:
// 		return ErrNotFound
// 	}
// }

// func main() {
// 	fmt.Println(GetPage("/passwd"))
// 	fmt.Println(GetPage("/user/1"))
// 	fmt.Println(GetPage("users"))
// }

package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
)

type error interface {
	Error() string
}

type MyErrorType struct {
	Err     error
	Code    int
	Message string
	debug   string
}

func (e *MyErrorType) Error() string { // interface
	return fmt.Sprintln(e.Err, e.Code, e.Message, e.debug)

}

func MyErrorFunc(err error, code int, message string) error { // функция для удобства ввода значений в структуру

	return &MyErrorType{
		Err:     err,
		Code:    code,
		Message: message,
		debug:   string(debug.Stack()),
	}

}

/*
func dev(a, b int) int {

	defer func() {

		if v := recover(); v != nil {
			fmt.Println("Функция восстановлена по причине:", v)
			// var err MyErrorType
			// err.Err = errors.New("new err")
			// fmt.Println(MyErrorFunc(&err, 44, "myfunc"))

		}

	}()

	res := a / b

	return res
}
*/

var (
	temp   = errors.New("Причина остановки функции")
	errNew MyErrorType
)

func OpenFile(path string) error {

	defer func() {

		if v := recover(); v != nil {
			fmt.Println("Функция OpenFile была восстановлена")
			errNew.Err = temp
			errNew.Err = fmt.Errorf("%w: %v", errNew.Err, v)
			errNew.debug = string(debug.Stack())

		}

	}()

	_, err := os.OpenFile(path, os.O_RDWR, 0777)

	if err != nil {
		// return err
		panic("OpenFile - имеет ошибку")

	}
	return err

}

func main() {
	// fmt.Println("Задание №1")
	// fmt.Println(dev(9, 0))

	// fmt.Println("Задание №2")

	_ = OpenFile("data.jason")
	// fmt.Println(err)
	fmt.Println(errNew)
	// Bufer, err := os.OpenFile("data.json", os.O_RDWR, 0777)
	// if err != nil {

	// 	fmt.Println(MyErrorFunc(err, 404, "open file"))
	// }

	fmt.Println("Приложение работает дальше")
}
