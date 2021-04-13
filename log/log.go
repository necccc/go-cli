package log

import (
	"fmt"
	"os"
)

func Write(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
