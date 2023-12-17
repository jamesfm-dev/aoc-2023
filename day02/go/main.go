package main

import (
	"log"
	"main/part"
	"os"
)

func main() {
	debug := len(os.Args) == 3 && (os.Args[2] == "-d" || os.Args[2] == "--debug")

	if len(os.Args) < 2 {
		log.Fatal("error no input provided\n")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error failed to open file %s: %v\n", filename, err)
	}
	defer f.Close()

	part.One(f, debug)
}
