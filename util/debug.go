// debug is print with debuglevels
package util

import (
	"fmt"
)

var debuglvl int

// SetDebuglvl 1-10 5=info 10=debug
func SetDebuglvl(n int) {
	debuglvl = n
}

// Debugf prints if Debuglvl more then 10
func Debugf(format string, a ...any) {
	if debuglvl > 9 {
		fmt.Printf(format, a...)
	}
}

// Debugf prints if Debuglvl more then 5
func Infof(format string, a ...any) {
	if debuglvl > 4 {
		fmt.Printf(format, a...)
	}
}
