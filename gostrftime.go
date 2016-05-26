// Copyright (c) 2014-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package gostrftime formats time.Time using strftime(3) conventions.
package gostrftime

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func strftime(b *bytes.Buffer, c rune, t time.Time) error {
	switch c {
	case 'A':
		b.WriteString(t.Weekday().String())
	case 'a':
		b.WriteString(t.Weekday().String()[:3])
	case 'B':
		b.WriteString(t.Month().String())
	case 'b':
		b.WriteString(t.Month().String()[:3])
	case 'C':
		fmt.Fprintf(b, "%02d", t.Year())
	case 'D':
		y, m, d := t.Date()
		fmt.Fprintf(b, "%02d/%02d/%02d", m, d, y%100)
	case 'd':
		fmt.Fprintf(b, "%02d", t.Day())
	case 'e':
		fmt.Fprintf(b, "%2d", t.Day())
	case 'F':
		y, m, d := t.Date()
		fmt.Fprintf(b, "%04d-%02d-%02d", y, m, d)
	case 'f':
		usec := t.Nanosecond() / 1e3
		fmt.Fprintf(b, "%06d", int(usec))
	case 'H':
		fmt.Fprintf(b, "%02d", t.Hour())
	case 'h':
		b.WriteString(t.Month().String()[:3])
	case 'I':
		hr := t.Hour() % 12
		if hr == 0 {
			hr = 12
		}
		fmt.Fprintf(b, "%02d", hr)
	case 'j':
		fmt.Fprintf(b, "%03d", t.YearDay())
	case 'k':
		fmt.Fprintf(b, "%2d", t.Hour())
	case 'L':
		msec := t.Nanosecond() / 1e6
		fmt.Fprintf(b, "%03d", int(msec))
	case 'l':
		hr := t.Hour() % 12
		if hr == 0 {
			hr = 12
		}
		fmt.Fprintf(b, "%2d", hr)
	case 'M':
		fmt.Fprintf(b, "%02d", t.Minute())
	case 'm':
		fmt.Fprintf(b, "%02d", t.Month())
	case 'N':
		fmt.Fprintf(b, "%09d", t.Nanosecond())
	case 'n':
		b.WriteByte('\n')
	case 'P':
		if t.Hour() < 12 {
			b.WriteString("am")
		} else {
			b.WriteString("pm")
		}
	case 'p':
		if t.Hour() < 12 {
			b.WriteString("AM")
		} else {
			b.WriteString("PM")
		}
	case 'R':
		h, m, _ := t.Clock()
		fmt.Fprintf(b, "%02d:%02d", h, m)
	case 'r':
		h, m, s := t.Clock()

		var tm string
		if h < 12 {
			tm = "AM"
		} else {
			tm = "PM"
		}

		hr := h % 12
		if hr == 0 {
			hr = 12
		}

		fmt.Fprintf(b, "%02d:%02d:%02d %s", hr, m, s, tm)
	case 'S':
		fmt.Fprintf(b, "%02d", t.Second())
	case 's':
		b.WriteString(strconv.FormatInt(t.Unix(), 10))
	case 'T':
		h, m, s := t.Clock()
		fmt.Fprintf(b, "%02d:%02d:%02d", h, m, s)
	case 't':
		b.WriteByte('\t')
	case 'v':
		fmt.Fprintf(b, "%2d-%s-%04d", t.Day(), t.Month().String()[:3], t.Year())
	case 'w':
		fmt.Fprintf(b, "%d", t.Weekday())
	case 'Y':
		fmt.Fprintf(b, "%04d", t.Year())
	case 'y':
		fmt.Fprintf(b, "%02d", t.Year()%100)
	case 'Z':
		zone, _ := t.Zone()
		b.WriteString(zone)
	case 'z':
		_, offset := t.Zone()
		allMinutes := int(offset / 60)
		fmt.Fprintf(b, "%+03d%02d", int(allMinutes/60), int(allMinutes%60))
	default:
		return fmt.Errorf("No valid replacement")
	}

	return nil
}

// Format returns a textual representation of the time value
// formatting according to format. format supports most of the formatting
// methods defined in strftime(3), minus the GNU libc extensions and POSIX
// locale extensions.
//
// List of accepted format expansion values:
//  %A  full weekday name (Sunday)
//  %a  abbreviated weekday name (Sun)
//  %B  full month name (September)
//  %b  abbreviated month name (Sep)
//  %C  (year / 100) as number. Single digits are preceded by zero (20)
//  %D  equivalent to %m/%d/%y (09/21/14)
//  %d  day of month as number. Single digits are preceded by zero (21)
//  %e  day of month as number. Single digits are preceded by a blank (21)
//  %f  microsecond as a six digit decimal number, zero-padded on the left (001234)
//  %F  equivalent to %Y-%m-%d (2014-09-21)
//  %H  the hour (24 hour clock) as a number. Single digits are preceded by zero (15)
//  %h  same as %b
//  %I  the hour (12 hour clock) as a number. Single digits are preceded by zero (03)
//  %j  the day of the year as a decimal number. Single digits are preced by zeros (264)
//  %k  the hour (24 hour clock) as a number. Single digits are preceded by a blank (15)
//  %L  millisecond as a three digit decimal number, zero-padded on the left (001)
//  %l  the hour (12 hour clock) as a number. Single digits are preceded by blank ( 3)
//  %M  the minute as a decimal number. Single digits are preceded by a zero (32)
//  %m  the month as a decimal number. Single digits are preceded by a zero (09)
//  %N  nanosecond as a 9 digit decimal number, zero-padded on the left (001234567)
//  %n  a newline (\n)
//  %P  am or pm as appropriate
//  %p  AM or PM as appropriate
//  %R  equivalent to %H:%M
//  %r  equivalent to %I:%M:%S %p
//  %S  the second as a number. Single digits are preceded by a zero (05)
//  %s  the number of seconds since the Epoch, UTC
//  %T  equivalent to %H:%M:%S
//  %t  a tab (\t)
//  %v  equivalent to %e-%b-%Y
//  %w  the weekday (Sunday as first day of the week) as a number. (0)
//  %Y  the year with century as a number (2014)
//  %y  year without century as a number. Single digits are preceded by zero (14)
//  %Z  time zone name (UTC)
//  %z  the time zone offset from UTC (-0700)
func Format(format string, t time.Time) string {
	if !strings.Contains(format, "%") {
		return format
	}

	outBuf := &bytes.Buffer{}
	rr := strings.NewReader(format)
	for {
		r, _, err := rr.ReadRune()
		if err != nil {
			break
		}

		if r != '%' {
			outBuf.WriteRune(r)
			continue
		}

		nr, _, err := rr.ReadRune()
		if err != nil {
			// got a percent, but then end of string
			// just append % and finish
			outBuf.WriteByte('%')
			break
		}
		if nr == '%' {
			outBuf.WriteByte('%')
			continue
		}

		err = strftime(outBuf, nr, t)
		if err != nil {
			outBuf.WriteByte('%')
			outBuf.WriteRune(nr)
			continue
		}
	}
	return outBuf.String()
}

// Strftime is an alias for Format
func Strftime(format string, t time.Time) string {
	return Format(format, t)
}
