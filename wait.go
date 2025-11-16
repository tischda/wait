package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

// number of progress bar increments
var TICKS = len(bar)

// Dumb implementation of a progress bar...
// see https://en.wikipedia.org/wiki/Block_Elements for unicode block elements
var bar = []string{
	"\r[          ]   0%",
	"\r[\u2591         ]  10%",
	"\r[\u2591\u2591        ]  20%",
	"\r[\u2591\u2591\u2591       ]  30%",
	"\r[\u2591\u2591\u2591\u2591      ]  40%",
	"\r[\u2591\u2591\u2591\u2591\u2591     ]  50%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591    ]  60%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591   ]  70%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591  ]  80%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591 ]  90%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591] 100%",
}

// wait until duration elapsed or a key is pressed on terminal
func wait(duration time.Duration, quiet bool) {

	stop := make(chan struct{})
	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		fmt.Printf("Waiting for %s, press a key to continue ...\n", duration)
		go watchKeypress(stop)
	} else {
		fmt.Printf("Waiting for %s ...\n", duration)
	}

	if quiet {
		select {
		case <-time.After(duration):
			return
		case <-stop:
			return
		}
	}

	// progress bar mode
	interval := duration / time.Duration(TICKS)
	hide_cursor()
	defer show_cursor()

	for i := 0; i < TICKS; i++ {
		fmt.Print(bar[i])
		select {
		case <-time.After(interval):
			// proceed to next tick
		case <-stop:
			return
		}
	}
}

// watchKeypress waits for any input by switch stdin into 'raw' mode
// cf. https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter
func watchKeypress(stop chan struct{}) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("term.MakeRaw():", err)
		return
	}
	defer func() {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			fmt.Println("term.Restore():", err)
		}
	}()

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println("os.Stdin.Read():", err)
		return
	}
	select {
	case <-stop: // already closed
	default:
		close(stop)
	}
}
