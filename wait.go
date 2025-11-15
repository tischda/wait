package main

import (
	"fmt"
	"time"
)

// Dumb implementation of a progress bar...
// see https://en.wikipedia.org/wiki/Block_Elements for unicode block elements
var bar = []string{
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

// number of progress bar increments
var TICKS = len(bar)

func wait(duration time.Duration, quiet bool) {
	if !quiet {
		go show_progress(duration)
	}
	time.Sleep(duration)
	if !quiet {
		// no go routine here to make sure the '100%' is printed
		fmt.Println(bar[TICKS-1])
	}
}

// prints the progress bar as time is passing by
func show_progress(d time.Duration) {
	interval := d / time.Duration(TICKS)
	hide_cursor()
	for i := 0; i < TICKS; i++ {
		fmt.Print(bar[i])
		time.Sleep(interval)
	}
	show_cursor()
}
