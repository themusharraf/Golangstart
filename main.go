package main

import "fmt"

func main() {
	number1 := "hello"
	number := "world"
	fmt.Println(number1 + number)

	fmt.Println(number1 + number)
}

// go build main.go 2 lik ga o'tkazish
// ./main 2 lik code ni run qilish
// go run main.go
