package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(stdout, "Invalid args")
		return
	}
	be := "are"
	if !isAnagram(os.Args[1], os.Args[2]) {
		be = "are not"
	}
	fmt.Fprintf(stdout, "%v and %v %v anagram\n", os.Args[1], os.Args[2], be)
}

func isAnagram(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}
	sl1 := strings.Split(s1, "")
	sl2 := strings.Split(s2, "")
	sort.Strings(sl1)
	sort.Strings(sl2)
	return reflect.DeepEqual(sl1, sl2)
}
