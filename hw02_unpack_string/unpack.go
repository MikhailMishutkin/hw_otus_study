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
	var result, res1, res2 strings.Builder // инициализирую переменную встроенного типа билдер для хранения результатов распаковки(в том числе промежуточных)
	var d rune
	for i, r := range s { //проходим по строке
		var res string  // внутрицикловая переменная для построения результирующей строки
		k := rune(s[i]) // приводим к рунам для преобразований
		l := rune(s[0]) // первый элемент, который может вызвать ошибку
		switch {
		case unicode.IsDigit(l):
			log.Fatal(ErrInvalidString)
		case unicode.IsDigit(k) && unicode.IsDigit(d):
			log.Fatal(ErrInvalidString)
		case unicode.IsDigit(k) && k == '0':
			result = res2
		case unicode.IsDigit(k) && k == '1':
			result = res1
		case unicode.IsDigit(k):
			dig, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal()
			}
			result = res2                        // присваиваем предпредыдущий результат
			res = strings.Repeat(string(d), dig) // повторяем предыдущее буквенное значение dig раз, записываем в промежуточную переменную
		default:
			res = string(r) // записываем букву в промежуточную переменную
			res2 = res1
		}
		d = k
		result.WriteString(res)
		res2 = res1
		res1 = result
	}
	return result.String(), nil
}

/*


	if unicode.IsDigit(l) {
		log.Fatal(ErrInvalidString)
	} else if unicode.IsDigit(k) { // проверяем значение цифра или нет
		switch unicode.IsDigit(k) {
		case condition:

		}
		if unicode.IsDigit(k) = 0 {

		}
		dig, err := strconv.Atoi(string(s[i]))
		if err != nil {
			log.Fatal()
		}
		res = strings.Repeat(string(d), dig-1) // повторяем предыдущее буквенное значение dig-1 раз, записываем в промежуточную переменную
	} else {
		res = string(r) // записываем букву в промежуточную переменную
	}
	if unicode.IsDigit(k) && unicode.IsDigit(d) {
		log.Fatal(ErrInvalidString)
	} */
