package part

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Two(input *os.File, debug bool) {
	answer := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		game := strings.Fields(scanner.Text())
		games := game[2:]
		gameNumber := strings.Split(game[1], ":")[0]
		colorSet := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		sets := make([]map[string]int, len(games)/2)

		for i := 0; i < len(sets); i = i + 1 {
			color := games[2*i+1]
			valueStr := games[2*i]

			color = strings.TrimRightFunc(color, func(r rune) bool {
				return r == ';' || r == ','
			})

			value, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Fatalf("failed to convert value %s: %v\n", valueStr, err)
			}

			if colorSet[color] < value {
				colorSet[color] = value
			}
		}

		answer += colorSet["red"] * colorSet["green"] * colorSet["blue"]

		if debug {
			fmt.Printf("[Game %s] = %v\n", gameNumber, games)
			fmt.Printf("[Set] = %v\n\n", colorSet)
		}
	}

	if debug {
		fmt.Printf("Answer: %d\n", answer)
		return
	}
	fmt.Print(answer)
}
