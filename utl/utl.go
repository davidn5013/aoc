package utl

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// Helper functions

// Debug is fmt.Printf false or true
func Debug(b bool, format string, v ...any) {
	if b {
		fmt.Printf("DEBUG : "+format, v...)
	}
}

// PanicIf short for if b {panic()} don't for get to change the func name to most
func PanicIf(b bool, format string, v ...any) {
	if b {
		fmt.Printf("PANIC : "+format, v...)
		panic("Using PanicIf()")
	}
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

// SetTimer Set a timmer and return a func
// that returns time.Duration from the timer
// func main() {
// stopTimer := SetTimer()
// ...
// fmt.Printf("Elapsed time:%v\n".stopTimer())
func SetTimer() func() time.Duration {
	t := time.Now()
	return func() time.Duration {
		return time.Since(t)
	}
}
