package main

import (
	"flag"
	"log"
	"main/part"
	"os"
)

func main() {
	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	debug := flagSet.Bool("debug", false, "print debug messages")
	parts := flagSet.Int("part", 1, "which part to run")
	flagSet.Parse(os.Args[2:])

	if len(os.Args) < 2 {
		log.Fatal("error no input provided\n")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error failed to open file %s: %v\n", filename, err)
	}
	defer f.Close()

	switch *parts {
	default:
		log.Fatalf("error part \"%d\" does not exist: available parts are [1, 2]", *parts)
	case 1:
		part.One(f, *debug)
	case 2:
		part.Two(f, *debug)
	}
}
