package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func cut(inputLines []string, fields []int, delimiter string, separated bool) {
	for _, line := range inputLines {
		if separated && !strings.Contains(line, delimiter) {
			continue
		}
		parts := strings.Split(line, delimiter)
		var selectedParts []string
		for _, field := range fields {
			selectedParts = append(selectedParts, parts[field-1])
		}
		fmt.Println(strings.Join(selectedParts, delimiter))
	}
}

func main() {
	fieldsStr := flag.String("f", "", "select fields (columns)")
	delimiter := flag.String("d", "\t", "use DELIM instead of TAB as the field delimiter")
	separated := flag.Bool("s", false, "only print lines containing delimiter")
	flag.Parse()

	var fields []int
	for _, fieldStr := range strings.Split(*fieldsStr, ",") {
		var field int
		fmt.Sscanf(fieldStr, "%d", &field)
		fields = append(fields, field)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	cut(inputLines, fields, *delimiter, *separated)
}
