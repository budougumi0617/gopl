//Prints its command-line argments
package ch1ex1
import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
	for i:= 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
