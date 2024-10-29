package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type cutOptions struct {
	fields    string
	delimiter string
	separated bool
}

func main() {
	var opts cutOptions
	flag.StringVar(&opts.fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&opts.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&opts.separated, "s", false, "только строки с разделителем")
	flag.Parse()

	if opts.fields == "" {
		fmt.Fprintln(os.Stderr, "Usage: cutUtil -f fields [-d delimiter] [-s] [file...]")
		os.Exit(1)
	}

	fields := parseFields(opts.fields)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, opts.delimiter)
		if opts.separated && len(columns) < 2 {
			continue
		}
		output := extractFields(columns, fields)
		if len(output) > 0 {
			fmt.Println(strings.Join(output, opts.delimiter))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func parseFields(fields string) []int {
	var result []int
	for _, field := range strings.Split(fields, ",") {
		var num int
		fmt.Sscanf(field, "%d", &num)
		result = append(result, num-1) // индексы начинаются с 0
	}
	return result
}

func extractFields(columns []string, fields []int) []string {
	var result []string
	for _, field := range fields {
		if field >= 0 && field < len(columns) {
			result = append(result, columns[field])
		}
	}
	return result
}
