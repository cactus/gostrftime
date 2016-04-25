gostrftime
==========

[![Build Status](https://travis-ci.org/cactus/gostrftime.png?branch=master)][1]
[![GoDoc](https://godoc.org/github.com/cactus/gostrftime?status.png)][2]

## About

A Go pkg for formatting time.Time in an strftime(3) like way.
Basically, strftime for Go, with a couple of additions.


## Installing

    $ go get github.com/cactus/gostrftime

## Supported formats:

| code | Description |
| ---- | --- |
| `%A` | replaced by full weekday name (Sunday) |
| `%a` | replaced by abbreviated weekday name (Sun) |
| `%B` | replaced by full month name (September) |
| `%b` | replaced by abbreviated month name (Sep) |
| `%C` | replaced by (year / 100) as number. Single digits are preceded by zero (20) |
| `%D` | equivalent to `%m/%d/%y` (09/21/14) |
| `%d` | replaced by day of month as number. Single digits are preceded by zero (21) |
| `%e` | replaced by day of month as number. Signle digits are preceded by a blank (21) |
| `%f` | replaced by microsecond as a six digit decimal number, zero-padded on the left (001234) |
| `%F` | equivalent to` %Y-%m-%d` (2014-09-21) |
| `%H` | replaced by the hour (24 hour clock) as a number. Single digits are preceded by zero (15) |
| `%h` | same as `%b` |
| `%I` | replaced by the hour (12 hour clock) as a number. Single digits are preceded by zero (03) |
| `%j` | replaced by the day of the year as a decimal number. Single digits are preced by zeros (26 |4)
| `%k` | replaced by the hour (24 hour clock) as a number. Single digits are preceded by a blank (15 |)
| `%L` | replaced by millisecond as a three digit decimal number, zero-padded on the left (001) |
| `%l` | replaced by the hour (12 hour clock) as a number. Single digits are preceded by blank ( 3) |
| `%M` | replaced by the minute as a decimal number. Single digits are preceded by a zero (32) |
| `%m` | replaced by the month as a decimal number. Single digits are preceded by a zero (09) |
| `%N` | replaced by nanosecond as a 9 digit decimal number, zero-padded on the left (001234567) |
| `%n` | replaced by a newline (`\n`) |
| `%P` | replaced by am or pm as appropriate |
| `%p` | replaced by AM or PM as appropriate |
| `%R` | equivalent to `%H:%M` |
| `%r` | equivalent to `%I:%M:%S %p` |
| `%S` | replaced by the second as a number. Single digits are preceded by a zero (05) |
| `%s` | replaced by the number of seconds since the Epoch, UTC |
| `%T` | equivalant to `%H:%M:%S` |
| `%t` | replaced by a tab (`\t`) |
| `%v` | equivalent to `%e-%b-%Y` |
| `%w` | replaced by the weekday (Sunday as first day of the week) as a number. (0) |
| `%Y` | replaced by the year with century as a number (2014) |
| `%y` | replaced by year without century as a number. Single digits are preceded by zero (14) |
| `%Z` | replaced by time zone name (UTC) |
| `%z` | replaced by the time zone offset from UTC (-0700) |


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
