package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

// number of progress bar increments
var TICKS = len(bar)

// wait until duration elapsed or a key (line) pressed
func wait(duration time.Duration, quiet bool) {
	stop := make(chan struct{})
	go watchKeypress(stop)

	if quiet {
		select {
		case <-time.After(duration):
			return
		case <-stop:
			return
		}
	}

	// progress mode
	full := show_progress(duration, stop)

	if full {
		// ensure final 100% with newline
		fmt.Println(bar[TICKS-1])
	} else {
		// interrupted: move to next line
		fmt.Println()
	}
}

// prints the progress bar as time is passing by; stops early if interrupted.
// done sends true if full duration elapsed, false if interrupted.
func show_progress(d time.Duration, stop <-chan struct{}) bool {
	interval := d / time.Duration(TICKS)
	hide_cursor()
	defer show_cursor()
	for i := 0; i < TICKS; i++ {
		select {
		case <-stop:
			return false
		default:
			fmt.Print(bar[i])
			time.Sleep(interval)
		}
	}
	return true
}

// watchKeypress waits for any input by switch stdin into 'raw' mode
// cf. https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter
func watchKeypress(stop chan struct{}) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	select {
	case <-stop: // already closed
	default:
		close(stop)
	}
}
