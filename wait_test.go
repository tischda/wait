package main

import (
	"log"
	"testing"
	"time"
)

// acceptable delta for the time effectively slept
const DELTA = 30 * time.Millisecond

func TestWait(t *testing.T) {
	params := []string{"0.9s", "0.900s", "900ms", "900000us", "900000000ns"}
	cfg := &Config{
		quiet:      true,
		noprogress: true,
	}

	for _, duration := range params {
		start := time.Now()
		timeDuration, err := time.ParseDuration(duration)
		if err != nil {
			log.Fatalln(err)
		}
		wait(timeDuration, cfg)
		stop := time.Now()

		actual := stop.Sub(start)
		expected, err := time.ParseDuration("0.9s")
		if err != nil {
			log.Fatalln(err)
		}
		if actual < expected-DELTA || actual > expected+DELTA {
			t.Errorf("Duration: %s, Expected: %s, but was: %s", duration, expected, actual)
		}
	}
}

func TestProgress(t *testing.T) {
	cfg := &Config{
		quiet: false,
	}
	duration := "1s"
	start := time.Now()
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
	}
	wait(timeDuration, cfg)
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
