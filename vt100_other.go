//go:build !windows

package main

func enableVirtualTerminalProcessing() (func(), error) {
	return func() {}, nil
}
