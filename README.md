gostrftime
==========

[![Build Status](https://travis-ci.org/cactus/gostrftime.png?branch=master)][1]
[![GoDoc](https://godoc.org/github.com/cactus/gostrftime?status.png)][2]

## About

Simple Go pkg for formatting time.Time in an strftime(3) like way.


## Installing

    $ go get github.com/cactus/gostrftime


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
