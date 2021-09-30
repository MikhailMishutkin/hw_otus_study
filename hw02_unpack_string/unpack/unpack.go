package unpack

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	var result strings.Builder // инициализирую переменную встроенного типа билдер для хранения результата распаковки
	var d rune
	for i, r := range s { //проходим по строке
		var res string

		k := rune(s[i])
		l := rune(s[0])
		if unicode.IsDigit(l) == true {
			log.Fatal(ErrInvalidString)
		}
		if unicode.IsDigit(k) == true { // проверяем значение цифра или нет
			dig, err := strconv.Atoi(string(s[i]))
			if err != nil {
				errors.New("not a digit")
			}
			res = strings.Repeat(string(d), dig-1) // повторяем предыдущее буквенное значение dig-1 раз, записываем в промежуточную переменную
		} else {
			res = string(r) // записываем букву в промежуточную переменную
		}
		if unicode.IsDigit(k) == true && unicode.IsDigit(d) == true {
			log.Fatal(ErrInvalidString)
		}
		d = k

		result.WriteString(res)
	}

	// Place your code here.
	return result.String(), nil
}
