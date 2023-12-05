// debug is print debug with lvl for debug or info
package util

import "fmt"

var debuglvl int

func SetDebuglvl(n int) {
	debuglvl = n
}

func Debugf(format string, a ...any) {
	if debuglvl > 9 {
		fmt.Printf(format, a...)
	}
}

func Infof(format string, a ...any) {
	if debuglvl > 4 {
		fmt.Printf(format, a...)
	}
}
