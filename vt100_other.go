//go:build !windows

package main

func enableVirtualTerminalProcessing() func() {
	return func() {}
}
