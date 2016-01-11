//Package ex01 Customize echo
package ex01

import (
	"fmt"
	"os"
)

//Prints its command-line argments
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
