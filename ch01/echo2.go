// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // loop through all args is a quadratic process
		s += sep + arg // old s will be garbage collected, could be costly
		sep = " "
	}
	fmt.Println(s)
}
