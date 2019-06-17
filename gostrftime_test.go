// Copyright (c) 2014-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

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
	{
		"t̢̙̻͖̥h̯͇̝%i͠s ̦̰̭ͅ%îí̲̬̟̗̞͈̥s̪͓ͅ ̵͉͕̲͎a̰ ̡͚̩t̪̺͍͇̫̜͘ȩ̩̠̟̗̮͔͚st̴͎ ͍̭̦͚̘̤̥%Y%̣͍̱̣̰̦̳͝m̧̞̥͚͉ ̧̝%͕̦̲Y҉̥̰̤̟͍̳ ̸͈̲̘̠̭͕%î̯͕̺̳͟ ̞̠̞̘̖͚̰x̻͜",
		"t̢̙̻͖̥h̯͇̝%i͠s ̦̰̭ͅ%îí̲̬̟̗̞͈̥s̪͓ͅ ̵͉͕̲͎a̰ ̡͚̩t̪̺͍͇̫̜͘ȩ̩̠̟̗̮͔͚st̴͎ ͍̭̦͚̘̤̥2009%̣͍̱̣̰̦̳͝m̧̞̥͚͉ ̧̝%͕̦̲Y҉̥̰̤̟͍̳ ̸͈̲̘̠̭͕%î̯͕̺̳͟ ̞̠̞̘̖͚̰x̻͜",
	},
	{
		"t̵̢̟̞͙̮͙̬̣̼̝͍͕̱̗͉̜̳̂͒̄̀́̕͘ḩ̧̮̬̪̘̱ͧͩ͗̿̌͋ͥ̂̾͗͑ͣͦ̂̋͢i̳͎̱͕͍͓̥̻̥̺͖̞̠̥̬͊̓̊́ͪ̎̽̇̍̚͞͝ͅs̩̥̳̞͑͂ͤ̾ͫ̈͛̈̒̋͒͟͝ ̛ͤ̊ͭͤͭ̉͌̄̍ͣͯͬͭͩͦ̈͆̈́͆͢҉̺̰͈̤͎͎̻̖̹ͅͅïͧ͌ͩ̌҉͙̰̹̻͇̪̻̯͓̻̰͉̤̠͕̲͜%Ysͣ͑͒͌ͯ͛͗ͦ͑ͭ̅̔̄͐̚͟͢҉̣̜̫͍͉̣̳̟̜͇̪̬̗̭̦̮͢ͅ ̷̡́ͬ̑ͩͮͪ̅͌ͯ͂ͥ̈̏̋͆͏̴̝̤̮̬̰͓͎͎͉ͅaͯ̓̒͒ͯ̐ͪ́́ͤ͑ͫ̇҉̨͟͝͏̗̝̭̲̯̭̳͖͔̪̱̘͔̼̱ ̨̻̻͔̠̘͉͔̹͔͊̀̉ͧ́̽ͤ̀̐́͐̑̀t̓̉͛͆ͫ̏ͧ̽ͨͧͦͦ̐ͣͪ̇͏̷̡̧̧͇̰͎̥͈͓̦̯̪̙͚̲̺̺̠ͅėͤ̔̾̄́ͯ̓ͤ͗ͣ̌̾҉͏̢̫̘͓͎̼̺̟̭͚͚͘͜s̲̖̥̳̯̣̣̣͎̲̫͖̤͔̝̼̼ͮ̒ͯͯ̈̈ͧ̃ͤͥ̀͘͢t̴̸̯̲̼͈͌̒̔̈́͑͐ͧͮ̀́ͧ̂͛ͯͭ̕͠ ͍̼͖̣͐̒ͧ̉ͭ͊̋̂̓̀͡%̪͍͓̪͇̭͙̹͂̑ͣ̑͒ͦ͢͜͝͠m̶̡̨̛̹̩̜̺͚͙̱̭͈̹̯̰͖̓ͣ̆͆̓ͬͦ͂͐̕ͅ ̸̧̮̘͖͈͙̲̝̟͙̱͕͓̖̞̰̈́̆͋ͩ̍̏ͦ̇̈̒̒̾̌̓͞ͅ%̶̒̀ͯ̔͆ͨͮ͂͌̏ͪ̏͆̒ͥ͌ͩ́͏̺̫̮͎̖͎̦̖͕͉̭̠̠̗̥̫ͅỲ̝̟̪̙̹͗̏̓͛͋̏͛̀͝ ̵̗͇̘̦̥͔͇̹͖̦̳̿̌̊̿̈́ͩͦ̆̓͆͌͊̍̿͘%̷̡̢̯̙͍͙̱̟̺̽̅̐̏́̇î̢͑ͫ̀͊̾̉̿ͣ̆̆̎̉ͨ̑̿͋̚҉̶̭̩̟̬͎̦̠͟͝ ̟̩̗͖̺͕͍̄̿ͥͨ̽̀x̸̧̨͍͍̠̟̺̟͙ͨ̑̇ͣ̅̅̆̉͆̂͒͋̍ͤ́͡",
		"t̵̢̟̞͙̮͙̬̣̼̝͍͕̱̗͉̜̳̂͒̄̀́̕͘ḩ̧̮̬̪̘̱ͧͩ͗̿̌͋ͥ̂̾͗͑ͣͦ̂̋͢i̳͎̱͕͍͓̥̻̥̺͖̞̠̥̬͊̓̊́ͪ̎̽̇̍̚͞͝ͅs̩̥̳̞͑͂ͤ̾ͫ̈͛̈̒̋͒͟͝ ̛ͤ̊ͭͤͭ̉͌̄̍ͣͯͬͭͩͦ̈͆̈́͆͢҉̺̰͈̤͎͎̻̖̹ͅͅïͧ͌ͩ̌҉͙̰̹̻͇̪̻̯͓̻̰͉̤̠͕̲͜2009sͣ͑͒͌ͯ͛͗ͦ͑ͭ̅̔̄͐̚͟͢҉̣̜̫͍͉̣̳̟̜͇̪̬̗̭̦̮͢ͅ ̷̡́ͬ̑ͩͮͪ̅͌ͯ͂ͥ̈̏̋͆͏̴̝̤̮̬̰͓͎͎͉ͅaͯ̓̒͒ͯ̐ͪ́́ͤ͑ͫ̇҉̨͟͝͏̗̝̭̲̯̭̳͖͔̪̱̘͔̼̱ ̨̻̻͔̠̘͉͔̹͔͊̀̉ͧ́̽ͤ̀̐́͐̑̀t̓̉͛͆ͫ̏ͧ̽ͨͧͦͦ̐ͣͪ̇͏̷̡̧̧͇̰͎̥͈͓̦̯̪̙͚̲̺̺̠ͅėͤ̔̾̄́ͯ̓ͤ͗ͣ̌̾҉͏̢̫̘͓͎̼̺̟̭͚͚͘͜s̲̖̥̳̯̣̣̣͎̲̫͖̤͔̝̼̼ͮ̒ͯͯ̈̈ͧ̃ͤͥ̀͘͢t̴̸̯̲̼͈͌̒̔̈́͑͐ͧͮ̀́ͧ̂͛ͯͭ̕͠ ͍̼͖̣͐̒ͧ̉ͭ͊̋̂̓̀͡%̪͍͓̪͇̭͙̹͂̑ͣ̑͒ͦ͢͜͝͠m̶̡̨̛̹̩̜̺͚͙̱̭͈̹̯̰͖̓ͣ̆͆̓ͬͦ͂͐̕ͅ ̸̧̮̘͖͈͙̲̝̟͙̱͕͓̖̞̰̈́̆͋ͩ̍̏ͦ̇̈̒̒̾̌̓͞ͅ%̶̒̀ͯ̔͆ͨͮ͂͌̏ͪ̏͆̒ͥ͌ͩ́͏̺̫̮͎̖͎̦̖͕͉̭̠̠̗̥̫ͅỲ̝̟̪̙̹͗̏̓͛͋̏͛̀͝ ̵̗͇̘̦̥͔͇̹͖̦̳̿̌̊̿̈́ͩͦ̆̓͆͌͊̍̿͘%̷̡̢̯̙͍͙̱̟̺̽̅̐̏́̇î̢͑ͫ̀͊̾̉̿ͣ̆̆̎̉ͨ̑̿͋̚҉̶̭̩̟̬͎̦̠͟͝ ̟̩̗͖̺͕͍̄̿ͥͨ̽̀x̸̧̨͍͍̠̟̺̟͙ͨ̑̇ͣ̅̅̆̉͆̂͒͋̍ͤ́͡",
	},
	{
		"nope %\uFEFF \u2665%\u2665 \u2665 \U0001f389 %\U0001f602 %M",
		"nope %\uFEFF \u2665%\u2665 \u2665 \U0001f389 %\U0001f602 04",
	},
	{
		"test %\U00103737 %\uFFFDtest%\U0010FFFE%...",
		"test %\U00103737 %\uFFFDtest%\U0010FFFE%...",
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
	{"%f", "001234"},
	{"%H", "03"},
	{"%h", "Jan"},
	{"%I", "03"},
	{"%j", "002"},
	{"%k", " 3"},
	{"%L", "001"},
	{"%l", " 3"},
	{"%M", "04"},
	{"%m", "01"},
	{"%N", "001234567"},
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

	tm := time.Date(2009, time.January, 2, 3, 4, 0, 1234567, time.UTC)
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
	tm := time.Date(2009, time.January, 2, 3, 4, 0, 1234567, location)
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
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 123456, time.UTC)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Format("%A %a %B %b %C %D %d %e %F %f %H %h %I %j %k %L %l %M %m %N %n %p %R %r %S %s %T %t %v %w %Y %y %Z %z", tm)
	}
}

func BenchmarkStrfFormatSimple(b *testing.B) {
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 1234567, time.UTC)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Format(formattests[0].format, tm)
	}
}
