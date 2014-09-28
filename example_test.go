package gostrftime_test

import (
	"fmt"
	"time"

	"github.com/cactus/gostrftime"
)

func ExampleFormat() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(gostrftime.Format("%Y-%m-%d", t))
	// Output:
	// 2009-11-10
}
