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
	var result, result1, result2 strings.Builder // инициализирую переменную встроенного типа билдер для хранения результатов распаковки(в том числе промежуточных)
	var d rune
	for i, r := range s { // проходим по строке
		var res string  // внутрицикловая переменная для построения результирующей строки
		k := rune(s[i]) // приводим к рунам для преобразований
		l := rune(s[0]) // первый элемент, который может вызвать ошибку
		switch {
		case unicode.IsDigit(l):
			return "Неудача, ссорян", ErrInvalidString
		case unicode.IsDigit(k) && unicode.IsDigit(d):
			return "Неудача, ссорян", ErrInvalidString
		case unicode.IsDigit(k) && k == '0':
			result = result2
		case unicode.IsDigit(k) && k == '1':
			result = result1
		case unicode.IsDigit(k):
			dig, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal()
			}
			result = result2                     // присваиваем предпредыдущий результат
			res = strings.Repeat(string(d), dig) // повторяем предыдущее буквенное значение dig раз, записываем в промежуточную переменную
		default:
			res = string(r) // записываем букву в промежуточную переменную
			result2 = result1
		}
		d = k
		result.WriteString(res)
		result2 = result1
		result1 = result
	}
	return result.String(), nil
}
