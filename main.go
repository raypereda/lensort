package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// main is a utility that sort the standard input lines by length and then alphabetically.
func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: lensort")
		fmt.Println("Sorts standard input by line length and alphabetically")
		fmt.Println("Example: cat words.txt | lensort")
		os.Exit(1)
	}

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	byLengthThenAlphabetically := func(i, j int) bool {
		return len(lines[i]) < len(lines[j]) || len(lines[i]) == len(lines[j]) && lines[i] < lines[j]
	}

	sort.Slice(lines, byLengthThenAlphabetically)

	for _, l := range lines {
		fmt.Println(l)
	}
}
