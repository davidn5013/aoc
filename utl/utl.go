package utl

import (
	"fmt"
	"os"
	"runtime"
)

// Helper functions

// Debug is fmt.Printf false or true
func Debug(b bool, format string, v ...any) {
	if b {
		fmt.Printf("DEBUG : "+format, v...)
	}
}

// PanicIf short for if b {panic()} don't for get to change the func name to most
func PanicIf(b bool, s string) {
	panic(s)
}

// CurrFuncName print out func for use int fmt.Printf and debug texts
func CurrFuncName() string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	return runtime.FuncForPC(counter).Name()
}
