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
	quiet   bool
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.quiet, "q", false, "")
	flag.BoolVar(&cfg.quiet, "quiet", false, "suppress non-error output")
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
  -q, --quiet
          suppress non-error output
  -?, --help
          display this help message
  -v, --version
          print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` 3s
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

	duration := flag.Arg(0)

	// only numbers or dot, unit missing?
	exp := regexp.MustCompile(`^[\d.]+$`)
	if exp.FindString(duration) != "" {
		duration += "s" // seconds
	}
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
	}

	stop := make(chan struct{})
	go watchKeypress(stop)
	wait(timeDuration, cfg.quiet, stop)

}
