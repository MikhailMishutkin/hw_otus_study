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
	// инициализирую переменную встроенного типа
	// билдер для хранения результатов распаковки(в том числе промежуточных)
	var result, resultBefore, resultBeforeLast strings.Builder
	var d rune
	for i, r := range s { // проходим по строке
		// внутрицикловая переменная для построения результирующей строки
		var res string
		// приводим к рунам для преобразований
		k := rune(s[i])
		// первый элемент, который может вызвать ошибку
		l := rune(s[0])
		switch {
		case unicode.IsDigit(l):
			return "Неудача, ссорян", ErrInvalidString
		case unicode.IsDigit(k) && unicode.IsDigit(d):
			return "Неудача, ссорян", ErrInvalidString
		case unicode.IsDigit(k) && k == '0':
			result = resultBeforeLast
		case unicode.IsDigit(k) && k == '1':
			result = resultBefore
		case unicode.IsDigit(k):
			dig, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal()
			}
			// присваиваем предпредыдущий результат
			result = resultBeforeLast
			// повторяем предыдущее буквенное
			// значение dig раз, записываем в промежуточную переменную
			res = strings.Repeat(string(d), dig)
		default:
			// записываем букву в промежуточную переменную
			res = string(r)
			resultBeforeLast = resultBefore
		}
		d = k
		result.WriteString(res)
		resultBeforeLast = resultBefore
		resultBefore = result
	}
	return result.String(), nil
}
