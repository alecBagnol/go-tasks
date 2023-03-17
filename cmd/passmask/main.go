package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// regex pattern for passwords up to 64 chars long - I took 8 as minimum length cuz it's usually the minimum length for passwords
	regex_pattern := regexp.MustCompile(`(?i)(?:key|pw|pass(?:word)?|secret)(?:=)([a-zA-Z0-9!@#\$\^%&*()\-=_\+\\\[\]\{\}|;:,.'"<>\?~/]{8,64})`)

	content, _ := os.ReadFile("test.txt")
	input := string(content)

	matches := regex_pattern.FindAllStringSubmatchIndex(input, -1)
	input_masked := input
	offset := 0

	for _, match := range matches {
		fmt.Println(match)
		start, end := match[2]+offset, match[3]+offset
		input_masked = input_masked[:start] + "*****" + input_masked[end:]
		offset += 5 - (end - start)
	}

	fmt.Println(input_masked)
}
