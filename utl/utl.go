package utl

import "fmt"

// Helper functions
// Debug is fmt.Printf false or true
func Debug(b bool, format string, v ...any) {
	if b {
		fmt.Printf("DEBUG : "+format, v...)
	}
}

func PanicIf(b bool, s string) {
	panic(s)
}
