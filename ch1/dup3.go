// dup3 prints the text of each line that appears more than
// once in the. It reads from a list of named files.
package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	counts := make(map[string]int)
	
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}