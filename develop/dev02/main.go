package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	_, err := strconv.Atoi(s)
	if err == nil {
		return "", fmt.Errorf("Incorrect string")
	}

	sb := strings.Builder{}

	for i, v := range s {
		if unicode.IsDigit(v) {
			cnt, _ := strconv.Atoi(string(v))
			for j := 0; j < cnt - 1; j++ {
				sb.WriteByte(s[i - 1])
			}
		} else {
			sb.WriteRune(v)
		}
	}

	return sb.String(), nil
}

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
}
