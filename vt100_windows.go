//go:build windows

package main

import (
	"os"

	"golang.org/x/sys/windows"
)

// enable virtual terminal processing on windows
// see https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
func enableVirtualTerminalProcessing() func() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32
	if err := windows.GetConsoleMode(stdout, &originalMode); err != nil {
		// may not be a terminal, ignore error
		return func() {}
	}

	newMode := originalMode | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	if err := windows.SetConsoleMode(stdout, newMode); err != nil {
		// may not be a terminal, ignore error
		return func() {}
	}

	return func() {
		windows.SetConsoleMode(stdout, originalMode)
	}
}
