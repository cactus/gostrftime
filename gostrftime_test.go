package gostrftime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var formattests = []struct {
	format   string
	expected string
}{
	{
		"/this/is/a/test/%Y/%m/%d/test.log",
		"/this/is/a/test/2009/01/02/test.log",
	},
	{
		"/this/is/a/test/%Y%m%d/test.log",
		"/this/is/a/test/20090102/test.log",
	},
	{
		"/this/is/a/test/Ymd/test.log",
		"/this/is/a/test/Ymd/test.log",
	},
	{
		"/this/is/a/test/%%Y/test.log",
		"/this/is/a/test/%Y/test.log",
	},
	{
		"%/this/is/a/test/%9/test.log%",
		"%/this/is/a/test/%9/test.log%",
	},
	{
		"%ü-%Y-%m-%d-%%%m-%%%%m-ü-世%界",
		"%ü-2009-01-02-%01-%%m-ü-世%界",
	},
	{"%A", "Friday"},
	{"%a", "Fri"},
	{"%B", "January"},
	{"%b", "Jan"},
	{"%C", "2009"},
	{"%D", "01/02/09"},
	{"%d", "02"},
	{"%e", " 2"},
	{"%F", "2009-01-02"},
	{"%H", "03"},
	{"%h", "Jan"},
	{"%I", "03"},
	{"%j", "002"},
	{"%k", " 3"},
	{"%l", " 3"},
	{"%M", "04"},
	{"%m", "01"},
	{"%n", "\n"},
	{"%p", "AM"},
	{"%R", "03:04"},
	{"%r", "03:04:00 AM"},
	{"%S", "00"},
	{"%s", "1230865440"},
	{"%T", "03:04:00"},
	{"%t", "\t"},
	{"%v", " 2-Jan-2009"},
	{"%w", "5"},
	{"%Y", "2009"},
	{"%y", "09"},
	{"%Z", "UTC"},
	{"%z", "+0000"},
}

func TestStrfFormat(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tm := time.Date(2009, time.January, 2, 3, 4, 0, 0, time.UTC)
	for _, tt := range formattests {
		output := Format(tt.format, tm)
		assert.Equal(tt.expected, output, fmt.Sprintf("%s not right", tt.format))
	}
}

func TestStrfFormatNotUTC(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// use a timezone that doesn't do daylight savings
	location := time.FixedZone("Saskatchewan", -6*60*60)
	tm := time.Date(2009, time.January, 2, 3, 4, 0, 0, location)
	for _, tt := range formattests {
		var expected string
		switch tt.format {
		case "%Z":
			expected = "Saskatchewan"
		case "%z":
			expected = "-0600"
		case "%s":
			expected = "1230887040"
		default:
			expected = tt.expected
		}
		output := Format(tt.format, tm)
		assert.Equal(expected, output, fmt.Sprintf("%s not right", tt.format))
	}
}

func BenchmarkStrfFormatAll(b *testing.B) {
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Format("%A %a %B %b %C %D %d %e %F %H %h %I %j %k %l %M %m %n %p %R %r %S %s %T %t %v %w %Y %y %Z %z", tm)
	}
}

func BenchmarkStrfFormatSimple(b *testing.B) {
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Format(formattests[0].format, tm)
	}
}
