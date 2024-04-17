package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type flags struct {
	pattern string
	after int
	before int
	context int
	count bool
	ignoreCase bool
	invert bool
	fixed bool
	lineNumber bool
}

func parseFlags() *flags {
	var f flags

	flag.StringVar(&f.pattern,"pattern", "", "search pattern")
	flag.IntVar(&f.after, "A", 0, "print +N lines after match")
	flag.IntVar(&f.before, "B", 0, "print +N lines before match")
	flag.IntVar(&f.context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&f.count, "c", false, "print count of matching lines")
	flag.BoolVar(&f.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&f.invert, "v", false, "invert match")
	flag.BoolVar(&f.fixed, "F", false, "fixed string match")
	flag.BoolVar(&f.lineNumber, "n", false, "print line number")

	flag.Parse()

	return &f
}

func grep(inputLines []string, f *flags) {
	var re *regexp.Regexp
	if f.fixed {
		f.pattern = regexp.QuoteMeta(f.pattern)
	}
	if f.ignoreCase {
		re = regexp.MustCompile("(?i)" +f.pattern)
	} else {
		re = regexp.MustCompile(f.pattern)
	}

	var matchedIndexes []int

	for i, line := range inputLines {
		matched := re.MatchString(line)
		if (f.invert && !matched) || (!f.invert && matched) {
			matchedIndexes = append(matchedIndexes, i)
		}
	}

	for _, index := range matchedIndexes {
		start := index - f.before
		if start < 0 {
			start = 0
		}
		end := index + f.after + 1
		if end > len(inputLines) {
			end = len(inputLines)
		}

		for i := start; i < end; i++ {
			if f.count {
				continue
			}

			matched := re.MatchString(inputLines[i])
			if (f.invert && !matched) || (!f.invert && matched) {
				if f.lineNumber {
					fmt.Printf("%d:", i+1)
				}
				fmt.Println(inputLines[i])
			}
		}
	}

	if f.count {
		fmt.Println(len(matchedIndexes))
	}
}

func main() {
	flags := parseFlags()

	scanner := bufio.NewScanner(os.Stdin)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	grep(inputLines, flags)
}
