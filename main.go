package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	nobreak    bool
	noprogress bool
	quiet      bool
	help       bool
	version    bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.nobreak, "b", false, "")
	flag.BoolVar(&cfg.nobreak, "nobreak", false, "ignore key presses and wait specified time")
	flag.BoolVar(&cfg.noprogress, "np", false, "")
	flag.BoolVar(&cfg.noprogress, "noprogress", false, "do not display progress bar")
	flag.BoolVar(&cfg.quiet, "q", false, "")
	flag.BoolVar(&cfg.quiet, "quiet", false, "suppress non-error output (implies --noprogress)")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] duration

Waits for specified duration or until key pressed.

OPTIONS:
  -b, --nobreak
          ignore key presses and wait specified time
  -np, --noprogress
          do not display progress bar
  -q, --quiet
          suppress non-error output (implies --noprogress)
  -?, --help
          display this help message
  -v, --version
          print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` 3s
  Waiting for 3s, press a key to continue ...
  [░░░░░░░░░░] 100%`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if cfg.quiet {
		cfg.noprogress = true
	}

	duration := parseDuration(flag.Arg(0))
	wait(duration, cfg)

	if !cfg.noprogress {
		fmt.Printf("\n")
	}
}

// Adds "s" suffix if no time unit provided
func parseDuration(duration string) time.Duration {
	exp := regexp.MustCompile(`^[\d.]+$`)
	if exp.FindString(duration) != "" {
		duration += "s" // seconds
	}
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
	}
	return timeDuration
}
