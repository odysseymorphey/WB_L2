package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type flags struct {
	k int
	n bool
	r bool
	u bool
}

func parseFlags() *flags {
	f := flags{}

	flag.IntVar(&f.k, "k", -1, "column to sort")
	flag.BoolVar(&f.n, "n", false, "sort as a number value")
	flag.BoolVar(&f.r, "r", false, "reverse sort")
	flag.BoolVar(&f.u, "u", false, "sort only unique strings")
	flag.Parse()

	return &f
}

func readFile(path string) []string {
	var data []string
	
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		txt := sc.Text()
		data = append(data, txt)
	}

	return data
}

func reverse(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func makeUnique(data []string) []string {
	m := make(map[string]bool)
	newData := make([]string, 0)

	for _, v := range data {
		if _, ok := m[v]; !ok {
			m[v] = true
			newData = append(newData, v)
		}
	}

	return newData
}

func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func sortByColumn(lines []string, col int) []string {
	m := make(map[string]string)

	for i := 0; i < len(lines); i++ {
		splittedLine := strings.Split(lines[i], " ")
		m[splittedLine[col]] = lines[i]
	}

	keys := getKeys(m)
	sort.Strings(keys)

	sortedLines := make([]string, len(lines))
	for i, key := range keys {
		sortedLines[i] = m[key]
	}

	return sortedLines
}

func writeFile(data []string) {
	file, err := os.Create("sorted_file")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	for _, v := range data {
		fmt.Fprintln(file, v)
	}
}

func SortFile(path string) {
	f := parseFlags()

	data := readFile(path)

	if f.k >= 0 {
		sortByColumn(data, f.k)
	} else {
		sort.Strings(data)
	}

	if f.r {
		reverse(data)
	}

	if f.u {
		data = makeUnique(data)
	}

	writeFile(data)

	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	SortFile("./sample")
}