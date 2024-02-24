package main

import (
	"cut/cut"
	"flag"
	"fmt"
)

func main() {
	field := flag.String("f", "1 2", "prints out the specified field from each line")
	delimiter := flag.String("d", "	", "prints out the specified field from each line using a delimiter")
	flag.Parse()

	var filename string
	var text []byte

	if filename = flag.Arg(0); filename != "" {
		text = cut.OpenFile(filename)
	} else {
		fmt.Println("No file specified.")
	}

	fieldList, err := cut.GetFields(*field)
	if err != nil {
		fmt.Println("Invalid fields format")
	}

	if len(fieldList) > 0 && *delimiter == "	" {
		output := cut.CutSpecificField(text, fieldList)
		for _, field := range output {
			fmt.Println(field)
		}
	} else {
		output := cut.CutSpecificFieldWithDelimiter(text, fieldList, *delimiter)
		for _, field := range output {
			fmt.Println(field)
		}
	}
}
