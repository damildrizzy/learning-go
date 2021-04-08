package main

import (
	"fmt"
	"strings"
)


type MyStr string //alias for type string

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyStr("test").Uppercase())
}