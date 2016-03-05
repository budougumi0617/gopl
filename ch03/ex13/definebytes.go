package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

// Byte units
const (
	STEP = 1000 // Unit step
	KB
	MB = KB * STEP
	GB = MB * STEP
	TB = GB * STEP
	PB = TB * STEP
	EB = PB * STEP
	ZB = EB * STEP
	YB = ZB * STEP
)

func main() {
	fmt.Fprintf(stdout, "ZB/EB = %v\n", ZB/EB)
}
