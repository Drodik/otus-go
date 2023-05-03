package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputStr string) (string, error) {
	var result strings.Builder
	var repeatCount int
	var repeatVal string
	var prevVal rune
	var err error

	writing := false
	backslash := false

	for i, val := range inputStr {
		currentVal := string(val)
		prevIsNum := unicode.IsDigit(prevVal)
		currentIsNum := unicode.IsDigit(val)

		if i == 0 && currentIsNum {
			return "", ErrInvalidString
		}

		if prevIsNum && currentIsNum {
			return "", ErrInvalidString
		}

		if currentIsNum {
			repeatCount, err = strconv.Atoi(currentVal)

			if err != nil {
				return "", err
			}

			writing = true
		} else {
			writing = false

			if prevVal > 0 && !unicode.IsDigit(prevVal) {
				writing = true
				repeatCount = 1
			}

			if backslash {
				writing = false
			}
		}

		if writing {
			repeatVal = string(prevVal)

			if backslash {
				repeatVal = strings.Join([]string{"\\", repeatVal}, "")
				fmt.Println(repeatVal, repeatCount)
				backslash = false
			}

			result.WriteString(strings.Repeat(repeatVal, repeatCount))

			if currentVal == "\\" {
				backslash = true
			}
		}

		if i == len(inputStr)-1 {
			result.WriteString(strings.Repeat(currentVal, 1))
		}

		prevVal = val
	}

	return result.String(), nil
}
