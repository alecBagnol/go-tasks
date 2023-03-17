package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// regex pattern for passwords up to 64 chars long
	regex_pattern := regexp.MustCompile(`[a-zA-Z0-9!@#\$\^%&*()\-=_\+\\\[\]\{\}|;:,.'"<>\?~` + "`" + `]{1,64}`)
	// read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// replace pass with mask
	masked := regex_pattern.ReplaceAllString(input, "*****")
	// print
	fmt.Println(masked)
}
