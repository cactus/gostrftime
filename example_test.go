// Copyright (c) 2014-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package gostrftime

import (
	"fmt"
	"time"
)

func ExampleFormat() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(Strftime("%Y-%m-%d", t))
	// Output:
	// 2009-11-10
}
