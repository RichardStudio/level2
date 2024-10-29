package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortOptions struct {
	column  int
	numeric bool
	reverse bool
	unique  bool
}

func main() {
	var opts sortOptions
	flag.IntVar(&opts.column, "k", -1, "указать колонку для сортировки (начиная с 1)")
	flag.BoolVar(&opts.numeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&opts.reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&opts.unique, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "Usage: sortutil [options] inputfile outputfile")
		os.Exit(1)
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input file:", err)
		os.Exit(1)
	}

	if opts.unique {
		lines = unique(lines)
	}

	sort.Slice(lines, func(i, j int) bool {
		a, b := lines[i], lines[j]
		if opts.column > 0 {
			columnsA := strings.Fields(a)
			columnsB := strings.Fields(b)
			if opts.column-1 < len(columnsA) && opts.column-1 < len(columnsB) {
				a, b = columnsA[opts.column-1], columnsB[opts.column-1]
			}
		}
		if opts.numeric {
			numA, errA := strconv.Atoi(a)
			numB, errB := strconv.Atoi(b)
			if errA == nil && errB == nil {
				if opts.reverse {
					return numA > numB
				}
				return numA < numB
			}
		}
		if opts.reverse {
			return a > b
		}
		return a < b
	})

	if err := writeLines(outputFile, lines); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing output file:", err)
		os.Exit(1)
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func unique(lines []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}
	return result
}
