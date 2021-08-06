// dup2 prints the text of each line that appears more than
// once in the. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)	
	for input.Scan() {
		counts[input.Text()]++
	}
}