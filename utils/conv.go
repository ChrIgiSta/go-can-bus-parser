package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func InterfaceToString(i any) string {
	return fmt.Sprintf("%v", i)
}

func ComaSeperatedDecimalsToAscii(in string) string {
	var intArray []byte = make([]byte, 0)

	split := strings.Split(in, ",")

	for _, element := range split {
		num, _ := strconv.Atoi(element)
		intArray = append(intArray, byte(num))
	}
	return string(intArray)
}
