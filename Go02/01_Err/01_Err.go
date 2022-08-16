package main

import "fmt"

type MyError struct {
	code    int
	message string
}

func NewMyError(code int, message string) MyError {
	return MyError{code: code, message: message}
}

func (e MyError) Error() string {
	return fmt.Sprintf("%s %d", e.code, e.message)
}

var ErrorNotFound = NewMyError(404, "Not Found")

func main() {
	fmt.Println(ErrorNotFound)
}
