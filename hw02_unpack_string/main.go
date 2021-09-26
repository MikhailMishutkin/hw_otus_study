package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println("Введите символьную строку для распаковки")
	fmt.Scan(&s)

	hw02unpackstring.Unpack(s)

}
