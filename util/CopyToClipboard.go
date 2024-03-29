package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CopyToClipboard is for macOS
func CopyToClipboard(text string) error {
	// command := exec.Command("pbcopy")
	command := exec.Command("clip")
	command.Stdin = bytes.NewReader([]byte(text))

	if err := command.Start(); err != nil {
		return fmt.Errorf("error starting pbcopy command: %w", err)
	}

	err := command.Wait()
	if err != nil {
		return fmt.Errorf("error running pbcopy %w", err)
	}

	return nil
}
