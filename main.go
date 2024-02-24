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
	delimiter := flag.String("d", "	", "prints out the specified field from each line using a delimiter")
	flag.Parse()

	var filename string
	var text []byte

	if filename = flag.Arg(0); filename != "" {
		text = OpenFile(filename)
	} else {
		fmt.Println("No file specified.")
	}

	if *field > 0 && *delimiter == " " {
		output := cut.CutSpecificField(text, *field)
		for _, field := range output {
			fmt.Println(field)
		}
	} else {
		output := cut.CutSpecificFieldWithDelimiter(text, *field, *delimiter)
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
