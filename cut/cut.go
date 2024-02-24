package cut

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func GetFields(fields string) (output []int, err error) {
	var splittedFields []string
	var parsedField int64

	if strings.Contains(fields, ",") {
		splittedFields = strings.Split(fields, ",")
		for _, f := range splittedFields {
			parsedField, err = strconv.ParseInt(f, 0, 0)
			if err != nil {
				return nil, err
			}
			output = append(output, int(parsedField))
		}
	} else {
		splittedFields = strings.Split(fields, " ")
		for _, f := range splittedFields {
			parsedField, err := strconv.ParseInt(f, 0, 0)
			if err != nil {
				return nil, err
			}
			output = append(output, int(parsedField))
		}
	}
	return
}

func CutSpecificField(text []byte, fieldList []int) []string {
	reader := bytes.NewBuffer(text)
	var result []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Fields(line)

		if len(fieldList) > 1 {
			var elems []string
			for _, f := range fieldList {
				elems = append(elems, fields[f-1])
			}
			result = append(result, strings.Join(elems, "	"))
		} else {
			result = append(result, fields[fieldList[0]-1])
		}
	}
	return result
}

func CutSpecificFieldWithDelimiter(text []byte, fieldList []int, delim string) (output []string) {
	reader := bytes.NewBuffer(text)
	var result []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Split(line, delim)

		if len(fieldList) > 1 {
			var elems []string
			for _, f := range fieldList {
				elems = append(elems, fields[f-1])
			}
			result = append(result, strings.Join(elems, delim))
		} else {
			result = append(result, fields[fieldList[0]-1])
		}
	}
	return result
}
