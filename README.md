gostrftime
==========

[![Build Status](https://travis-ci.org/cactus/gostrftime.png?branch=master)][1]
[![GoDoc](https://godoc.org/github.com/cactus/gostrftime?status.png)][2]
[![Go Report Card](https://goreportcard.com/badge/cactus/gostrftime)](https://goreportcard.com/report/cactus/gostrftime)

## About

A Go pkg for formatting time.Time in an strftime(3) like way.
Basically, strftime for Go, with a couple of additions.


## Installing

    $ go get github.com/cactus/gostrftime

## Supported formats:

| code | example | Description |
| ---- | ------- | --- |
| `%A` | `Sunday` | full weekday name |
| `%a` | `Sun` | abbreviated weekday name |
| `%B` | `September` | full month name |
| `%b` | `Sep` | abbreviated month name |
| `%C` | `20` | (`year / 100`) as number. Single digits are preceded by zero |
| `%D` | `09/21/14` | equivalent to `%m/%d/%y` |
| `%d` | `21` | day of month as number. Single digits are preceded by zero |
| `%e` | `21` | day of month as number. Single digits are preceded by a blank |
| `%f` | `001234` | microsecond as a six digit decimal number, zero-padded on the left |
| `%F` | `2014-09-21` | equivalent to` %Y-%m-%d` |
| `%H` | `15` | the hour (24 hour clock) as a number. Single digits are preceded by zero |
| `%h` | `Sep` | same as `%b` |
| `%I` | `03` | the hour (12 hour clock) as a number. Single digits are preceded by zero |
| `%j` | `264` | the day of the year as a decimal number. Single digits are preced by zeros |
| `%k` | `15` | the hour (24 hour clock) as a number. Single digits are preceded by a blank |
| `%L` | `001` | millisecond as a three digit decimal number, zero-padded on the left |
| `%l` | `11` | replaced by the hour (12 hour clock) as a number. Single digits are preceded by blank |
| `%M` | `32` | replaced by the minute as a decimal number. Single digits are preceded by a zero |
| `%m` | `09` | replaced by the month as a decimal number. Single digits are preceded by a zero |
| `%N` | `001234567` | nanosecond as a 9 digit decimal number, zero-padded on the left |
| `%n` | `\n` | a newline |
| `%P` | `am` | am or pm as appropriate |
| `%p` | `AM` | AM or PM as appropriate |
| `%R` | `15:32` | equivalent to `%H:%M` |
| `%r` | `03:32:05 AM` | equivalent to `%I:%M:%S %p` |
| `%S` | `05` | the second as a number. Single digits are preceded by a zero |
| `%s` | `1461497457` | the number of seconds since the Epoch, UTC |
| `%T` | `15:32:05` | equivalant to `%H:%M:%S` |
| `%t` | `\t` | a tab |
| `%v` | `21-Sep-2014` | equivalent to `%e-%b-%Y` |
| `%w` | `0` |the weekday (Sunday as first day of the week) as a number |
| `%Y` | `2014` | replaced by the year with century as a number |
| `%y` | `14` | year without century as a number. Single digits are preceded by zero |
| `%Z` | `UTC` | time zone name |
| `%z` | `-0700` | the time zone offset from UTC |


## Using


    import (
        "fmt"
        "time"
        "github.com/cactus/gostrftime"
    )

    func main() {
        now := time.Now()
        fmt.Println(gostrftime.Format("%Y-%m-%d", now))
    }


## License

Released under an [MIT license][3]. See `LICENSE.md` file for details.

[1]: https://travis-ci.org/cactus/gostrftime
[2]: https://godoc.org/github.com/cactus/gostrftime
[3]: http://www.opensource.org/licenses/mit-license.php
