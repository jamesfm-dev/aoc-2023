package part

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func One(input *os.File, debug bool) {
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
		setCounter := 1
		valid := true

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

			if (color == "red" && value > 12) ||
				(color == "green" && value > 13) ||
				(color == "blue" && value > 14) {
				valid = false
			}

			colorSet[color] = value
			sets[setCounter-1] = colorSet
			if strings.ContainsRune(games[2*i+1], ';') {
				clear(colorSet)
				setCounter++
			}
		}

		if debug {
			fmt.Printf("[Game %s] = %v\n", gameNumber, games)
			fmt.Printf("[Sets %d] = %v\n\n", setCounter, sets[:setCounter])
		}

		if valid {
			x, err := strconv.Atoi(gameNumber)
			if err != nil {
				log.Fatalf("error failed to convert gameNum to int %s: %v", gameNumber, err)
			}
			answer += x
		}
	}

	if debug {
		fmt.Printf("Answer: %d\n", answer)
		return
	}
	fmt.Print(answer)
}
