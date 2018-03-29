// Print the names of all files passed as command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "filename: %v\n", err)
			continue
		}
		fmt.Println(f.Name()+":", "number of occurances")
	}
	fmt.Println("hello world")
}
