package goextra

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadNumbers(path string) []int {
	lines := ReadText(path)
	result := make([]int, len(lines))
	for i, line := range lines {
		number := StringToInt(line)
		result[i] = number
	}
	return result
}

func ReadText(path string) []string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Could not read file path ", path, " ", err)
	}

	lines := strings.Split(string(content), "\n")
	used_lines := make([]string, 0)
	for _, line := range lines {
		if line != "" {
			used_lines = append(used_lines, strings.TrimSuffix(line, "\r"))
		}
	}
	return used_lines
}
