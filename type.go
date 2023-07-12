package main

import (
	"fmt"
	"reflect"
)

func main() {
	n := 4
	m := 5
	fmt.Println(reflect.TypeOf(n + m))
}
