package main

import "fmt"

type HttpError struct {
	code    int
	message string
}

func (e HttpError) Error() string {
	return fmt.Sprintln(e.code, e.message)
}

func FuncHttpError(code int, message string) HttpError {
	return HttpError{code: code, message: message}
}

var (
	ErrNotFound  = FuncHttpError(404, "not found")
	ErrForbidden = FuncHttpError(403, "forbidden")
)

func GetPage(path string) error {
	switch path {
	case "/passwd":
		return ErrForbidden
	case "/user/1":
		return nil
	default:
		return ErrNotFound
	}
}

func main() {
	fmt.Println(GetPage("/passwd"))
	fmt.Println(GetPage("/user/1"))
	fmt.Println(GetPage("users"))
}
