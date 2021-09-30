package main

import (
	"fmt"
	"github.com/MikhailMishutkin/hw_otus_study/hw02_unpack_string"
)

//"hw_otus_study/hw02_unpack_string"
func main() {
	var s string
	fmt.Println("Введите символьную строку для распаковки")
	fmt.Scan(&s)
	a := hw02unpackstring.Unpack(s)
	fmt.Println(a)
}
