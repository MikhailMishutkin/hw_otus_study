package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/Quasilyte/concat"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	var result string
	for i, r := range s {
		var res string
		if unicode.IsDigit(r) == true {
			dig, err := strconv.Atoi(string(s[i]))
			if err != nil {
				errors.New("not a digit")
			}
			res = strings.Repeat(string(s[i-1]), dig)
		} else {
			res = string(r)
		}

		result = concat.Strings(result, res)
	}
	// Place your code here.
	return result, nil
}
