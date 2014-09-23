gostrftime
==========

[![GoDoc](https://godoc.org/github.com/cactus/gostrftime?status.png)][1]

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
        fmt.Println(gostrftime.Strftime("%Y-%m-%d", now))
    }


## License

Released under an [MIT license][2]. See `LICENSE.md` file for details.

[1]: https://godoc.org/github.com/cactus/gostrftime
[2]: http://www.opensource.org/licenses/mit-license.php
