package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error no input provided\n")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error failed to open file %s: %v\n", filename, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	answer := 0
	for scanner.Scan() {
		digits := []string{}
		for _, ch := range scanner.Text() {
			if '0' < ch && ch <= '9' {
				digits = append(digits, string(ch))
			}
		}

		if len(digits) <= 0 {
			continue
		}
		if len(digits) == 1 {
			digits = append(digits, digits[0])
		}

		s := strings.Join([]string{digits[0], digits[len(digits)-1]}, "")
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("failed to convert string to int %s: %v\n", s, err)
		}

		answer += i
	}

	fmt.Printf("%d", answer)
}
