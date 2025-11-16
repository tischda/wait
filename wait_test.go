package main

import (
	"log"
	"testing"
	"time"
)

// acceptable delta for the time effectively slept
const DELTA = 20 * time.Millisecond

func TestWait(t *testing.T) {
	params := []string{"0.1s", "0.100s", "100ms", "100000us", "100000000ns"}

	for _, duration := range params {
		start := time.Now()
		timeDuration, err := time.ParseDuration(duration)
		if err != nil {
			log.Fatalln(err)
		}
		wait(timeDuration, true)
		stop := time.Now()

		actual := stop.Sub(start)
		expected, err := time.ParseDuration("0.1s")
		if err != nil {
			log.Fatalln(err)
		}
		if actual < expected-DELTA || actual > expected+DELTA {
			t.Errorf("Duration: %s, Expected: %s, but was: %s", duration, expected, actual)
		}
	}
}

func TestProgress(t *testing.T) {
	duration := "1s"
	start := time.Now()
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
	}
	wait(timeDuration, false)
	stop := time.Now()

	actual := stop.Sub(start)
	expected, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
	}
	if actual < expected-DELTA || actual > expected+DELTA {
		t.Errorf("Duration: %s, Expected: %s, but was: %s", duration, expected, actual)
	}
}
