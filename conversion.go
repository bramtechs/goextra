package goextra

import (
	"log"
	"strconv"
)

func StringToInt(text string) int {
	result, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func BinaryStringToInt(text string) int {
	result, err := strconv.ParseInt(text, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(result)
}
