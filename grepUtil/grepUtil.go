package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Опции для утилиты
type grepOptions struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	var opts grepOptions
	flag.IntVar(&opts.after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&opts.before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&opts.context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&opts.count, "c", false, "количество строк")
	flag.BoolVar(&opts.ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&opts.invert, "v", false, "исключать совпадения")
	flag.BoolVar(&opts.fixed, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&opts.lineNum, "n", false, "напечатать номер строки")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintln(os.Stderr, "Usage: grepUtil [options] pattern filename")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	lines, err := readLines(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	matches := grep(lines, pattern, opts)
	for _, match := range matches {
		fmt.Println(match)
	}
}

// Чтение строк из файла
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

// grep функция для поиска строк
func grep(lines []string, pattern string, opts grepOptions) []string {
	var result []string
	var count int
	pattern = preparePattern(pattern, opts.ignoreCase)

	for i, line := range lines {
		lineToCheck := line
		if opts.ignoreCase {
			lineToCheck = strings.ToLower(line)
		}
		match := strings.Contains(lineToCheck, pattern)
		if opts.fixed {
			match = lineToCheck == pattern
		}
		if opts.invert {
			match = !match
		}
		if match {
			count++
			if opts.count {
				continue
			}
			result = appendContext(lines, result, i, opts)
		}
	}

	if opts.count {
		return []string{fmt.Sprintf("%d", count)}
	}

	return result
}

// Подготовка паттерна для игнорирования регистра
func preparePattern(pattern string, ignoreCase bool) string {
	if ignoreCase {
		return strings.ToLower(pattern)
	}
	return pattern
}

// Форматирование строки с учетом номера строки
func formatMatch(line string, lineNumber int, opts grepOptions) string {
	if opts.lineNum {
		return fmt.Sprintf("%d:%s", lineNumber+1, line)
	}
	return line
}

// Добавление контекста до и после совпадения
func appendContext(lines, result []string, i int, opts grepOptions) []string {
	before := opts.before
	after := opts.after

	if opts.context > 0 {
		before = opts.context
		after = opts.context
	}

	// Добавление строк до совпадения
	for j := i - before; j < i; j++ {
		if j >= 0 {
			result = append(result, lines[j])
		}
	}

	// Добавление строки совпадения
	if opts.lineNum {
		result = append(result, fmt.Sprintf("%d:%s", i+1, lines[i]))
	} else {
		result = append(result, lines[i])
	}

	// Добавление строк после совпадения
	for j := i + 1; j <= i+after && j < len(lines); j++ {
		result = append(result, lines[j])
	}

	return result
}
