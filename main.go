package main

import (
	"cut/cut"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	field := flag.Int("f", 1, "prints out the specified field from each line")
	flag.Parse()

	var filename string
	var text []byte

	if filename = flag.Arg(0); filename != "" {
		text = OpenFile(filename)
	} else {
		fmt.Println("No file specified.")
	}

	if *field > 0 {
		output := cut.CutSpecificField(text, *field)
		for _, field := range output {
			fmt.Println(field)
		}
	}
}

func OpenFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return text
}
