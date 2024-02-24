package cut

import (
	"bytes"
	"strings"
)

func CutSpecificField(text []byte, field int) []string {
	reader := bytes.NewBuffer(text)
	var result []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		result = append(result, fields[field-1])
	}
	return result
}
