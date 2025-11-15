[![Build Status](https://github.com/tischda/wait/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/wait/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/wait/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/wait/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/wait/badge.svg)](https://coveralls.io/r/tischda/wait)
[![Linter Status](https://github.com/tischda/wait/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/wait/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/wait)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/wait.svg)](https://github.com/tischda/wait/releases/latest)

Wait for specified duration or until key pressed.

## Install

~~~
go install github.com/tischda/wait
~~~

## Usage

~~~
Usage: wait [OPTIONS] duration

Waits for specified duration or until key pressed.

OPTIONS:
  -q, --quiet
          suppress non-error output
  -?, --help
          display this help message
  -v, --version
          print version and exit

EXAMPLES:

  $ wait 3s
    [░░░░░░░░░░] 100%
~~~

Will sleep for 3s (see [ParseDuration](http://golang.org/pkg/time/#ParseDuration) for time formats).

Accuracy for actual duration is not great, but the delta should be under 20ms.


## Inspiration

* https://github.com/tianon/gosleep
* https://github.com/cheggaaa/pb
* https://github.com/gosuri/uiprogress
