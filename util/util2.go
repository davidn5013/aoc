package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Helper functions

// Run execute a external command end program on fail
func Run(command string, arg ...string) {
	fmt.Printf("Running: %s with args %v\n", command, arg)
	cmd := exec.Command(command, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)
	}

	stdin := bufio.NewScanner(stdout)
	for stdin.Scan() {
		fmt.Println(stdin.Text())
	}
	cmd.Wait()
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
