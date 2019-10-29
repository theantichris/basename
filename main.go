// Basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	// Discard last '/' and everything before.
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	// Preserve everything before the last '.'.
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}

	// NOTE: ignoring potential errors from input.Err()
}
