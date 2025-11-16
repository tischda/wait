package main

import (
	"log"
	"testing"
	"time"
)

// acceptable delta for the time effectively slept
const DELTA_MS = 20 * time.Millisecond

func TestSleep(t *testing.T) {
	params := []string{"0.1s", "0.100s", "100ms", "100000us", "100000000ns"}

	for _, duration := range params {
		start := time.Now()
		timeDuration, err := time.ParseDuration(duration)
		if err != nil {
			log.Fatalln(err)
		}
		s := make(chan struct{})
		wait(timeDuration, true, s)
		stop := time.Now()

		actual := stop.Sub(start)
		expected, err := time.ParseDuration("0.1s")
		if err != nil {
			log.Fatalln(err)
		}
		if actual < expected-DELTA_MS || actual > expected+DELTA_MS {
			t.Errorf("Duration: %s, Expected: %s, but was: %s", duration, expected, actual)
		}
	}
}
